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

	context := context.Background()
	currentUser, err := state.db.GetUser(context, state.config.CurrentUserName)

	if err != nil {
		return fmt.Errorf("command: addFeed cannot get current user %w", err)
	}

	feed, errFeed := state.db.CreateFeed(context, database.CreateFeedParams{
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

	fmt.Println("Feed created successfully!")
	printFeed(feed)
	return nil
}

func printFeed(feed database.Feed) {
	fmt.Printf(" * ID: %s\n", feed.ID)
	fmt.Printf(" * Name: %s\n", feed.Name)
	fmt.Printf(" * URL: %s\n", feed.Url)
}
