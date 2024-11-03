package main

import (
	"context"
	"fmt"

	"github.com/trantuvan/bootdev-gator/internal/rssfeed"
)

func handlerAgg(state *state, command command) error {
	feed, err := rssfeed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")

	if err != nil {
		return fmt.Errorf("failed to fetch RssFeed %w", err)
	}

	fmt.Printf("%+v\n", *feed)
	return nil
}
