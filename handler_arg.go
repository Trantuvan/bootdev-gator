package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/trantuvan/bootdev-gator/internal/database"
	"github.com/trantuvan/bootdev-gator/internal/rssfeed"
)

var ErrDupUrlKey = errors.New("pq: duplicate key value violates unique constraint")

func handlerAgg(state *state, command command) error {
	if len(command.args) < 1 || len(command.args) >= 2 {
		return fmt.Errorf("command %s: takes 1 <time_between_reqs>", command.name)
	}

	timeBetweenRequests, err := time.ParseDuration(command.args[0])

	if err != nil {
		return fmt.Errorf("command %s: duration wrong format %w", command.name, err)
	}

	fmt.Printf("Collecting feeds every %s...", timeBetweenRequests)

	for range time.Tick(timeBetweenRequests) {
		scrapeFeeds(state)
	}

	return nil
}

func scrapeFeeds(state *state) error {
	ctx := context.Background()
	nextFeed, err := state.db.GetNextFeedToFetch(ctx)

	if err != nil {
		return fmt.Errorf("scrapeFeeds: cannot fetch next feed %w", err)
	}

	_, err = state.db.MarkFeedFetched(ctx, nextFeed.ID)

	if err != nil {
		return fmt.Errorf("scrapeFeeds: cannot mark feed is fetched %w", err)
	}

	feedData, err := rssfeed.FetchFeed(ctx, nextFeed.Url)

	if err != nil {
		return fmt.Errorf("scrapeFeeds: cannot get feed from url:%s -> %w", nextFeed.Url, err)
	}

	for _, item := range feedData.Channel.Item {
		publishedAt, errPub := time.Parse(time.UnixDate, item.PubDate)

		_, err := state.db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: publishedAt, Valid: errPub != nil},
			FeedID:      nextFeed.ID,
		})

		if err != nil && err != ErrDupUrlKey {
			log.Printf("scrapeFeeds: cannot create post from url:%s -> %v", nextFeed.Url, err)
		}
	}

	fmt.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(feedData.Channel.Item))
	return nil
}
