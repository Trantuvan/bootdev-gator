package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/trantuvan/bootdev-gator/internal/database"
)

func handlerFollow(state *state, command command) error {
	if len(command.args) == 0 {
		return fmt.Errorf("command: follow expected 1 arg url")
	}

	ctx := context.Background()
	currentUser, err := state.db.GetUser(ctx, state.config.CurrentUserName)

	if err != nil {
		return fmt.Errorf("command: follow cannot get current user %w", err)
	}

	followFeed, err := state.db.GetFeed(ctx, command.args[0])

	if err != nil {
		return fmt.Errorf("command: follow cannot get follow feed %w", err)
	}

	createFeedFollowed, err := state.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    currentUser.ID,
		FeedID:    followFeed.ID,
	})

	if err != nil {
		return fmt.Errorf("command: follow cannot create follow feed %w", err)
	}

	printFeedFollow(createFeedFollowed)
	return nil
}

func handlerFollowing(state *state, command command) error {
	feeds, err := state.db.GetFeedFollowsForUser(context.Background(), state.config.CurrentUserName)

	if err != nil {
		return fmt.Errorf("command: following cannot get feeds currentUser %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf(" * Feed Name: %s\n", feed.FeedName)
	}
	fmt.Printf(" * Current User: %s\n", state.config.CurrentUserName)

	return nil
}

func printFeedFollow(feed database.CreateFeedFollowRow) {
	fmt.Printf(" * Feed Name: %s\n", feed.FeedName)
	fmt.Printf(" * Current User: %s\n", feed.UserName)
}
