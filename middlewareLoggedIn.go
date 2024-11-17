package main

import (
	"context"
	"fmt"

	"github.com/trantuvan/bootdev-gator/internal/database"
)

func middlewareLoggedIn(handler func(state *state, command command, user database.User) error) func(*state, command) error {
	//* return func to register (stop handler exec right away)
	return func(state *state, command command) error {
		//* 1. before exec: setup loggedIn
		user, err := state.db.GetUser(context.Background(), state.config.CurrentUserName)

		if err != nil {
			return fmt.Errorf("middleware logged in failed: %w", err)
		}

		//* 2. exec handler, (handler value is error)
		return handler(state, command, user)
	}
}
