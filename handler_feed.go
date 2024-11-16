package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/trantuvan/bootdev-gator/internal/database"
)

func handlerAddFeed(state *state, command command) error {
	if len(command.args) < 2 {
		return fmt.Errorf("command: addFeed expected 2 args name url")
	}

	ctx := context.Background()
	currentUser, err := state.db.GetUser(ctx, state.config.CurrentUserName)

	if err != nil {
		return fmt.Errorf("command: addFeed cannot get current user %w", err)
	}

	feed, errFeed := state.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      command.args[0],
		Url:       command.args[1],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
	})

	if errFeed != nil {
		return fmt.Errorf("command: addFeed cannot create new feed %w", errFeed)
	}

	_, err = state.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("command: addFeed cannot follow user created feed %w", err)
	}

	fmt.Println("Feed created successfully!")
	printFeed(feed)
	return nil
}

func handlerFeeds(state *state, command command) error {
	feedUsers, err := state.db.GetFeeds(context.Background())

	if err != nil {
		return fmt.Errorf("command: failed to get all feeds %w", err)
	}

	for _, feedUser := range feedUsers {
		fmt.Printf(" * FeedName: %s\n", feedUser.FeedName)
		fmt.Printf(" * FeedUrl: %s\n", feedUser.Url)
		fmt.Printf(" * UserName: %s\n", feedUser.UserName)
	}
	return nil
}
func printFeed(feed database.Feed) {
	fmt.Printf(" * ID: %s\n", feed.ID)
	fmt.Printf(" * Name: %s\n", feed.Name)
	fmt.Printf(" * URL: %s\n", feed.Url)
}
