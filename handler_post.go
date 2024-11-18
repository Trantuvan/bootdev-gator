package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/trantuvan/bootdev-gator/internal/database"
)

func handlerBrowse(state *state, command command, user database.User) error {
	take, skip := 2, 2

	if len(command.args) == 1 {
		take, _ = strconv.Atoi(command.args[0])
	} else if len(command.args) == 2 {
		take, _ = strconv.Atoi(command.args[0])
		skip, _ = strconv.Atoi(command.args[1])
	} else {
		fmt.Printf("command %s: takes 2 optinal [<take>] [<skip>]\n", command.name)
	}

	posts, err := state.db.GetPostsByUser(context.Background(), database.GetPostsByUserParams{
		ID:     user.ID,
		Limit:  int32(take),
		Offset: int32(skip),
	})

	if err != nil {
		return fmt.Errorf("command %s: cannot get posts -> %w", command.name, err)
	}

	for _, post := range posts {
		fmt.Printf("* Title: %s\n", post.Title)
		fmt.Printf("* PubDate: %v+\n", post.PublishedAt)
		fmt.Printf("* Desc: %v\n", post.Description)
	}
	return nil
}
