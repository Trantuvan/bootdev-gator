package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/trantuvan/bootdev-gator/internal/database"
)

func handlerLogin(state *state, command command) error {
	if len(command.args) == 0 {
		return fmt.Errorf("command: login expected 1 arg username")
	}

	exiestedUser, err := state.db.GetUser(context.Background(), command.args[0])

	if err != nil {
		os.Exit(1)
	}

	if err := state.config.SetUser(exiestedUser.Name); err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}

	fmt.Printf("username: %s has been set", state.config.CurrentUserName)
	return nil
}

func handlerRegister(state *state, command command) error {
	if len(command.args) == 0 {
		return fmt.Errorf("command: register expected 1 arg username")
	}

	createdUser, err := state.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      command.args[0],
	})

	if err != nil {
		return fmt.Errorf("faild to register new user: %w", err)
	}

	state.config.SetUser(createdUser.Name)
	fmt.Println("User created successfully!")
	printUser(createdUser)
	return nil
}

func handlerUsers(state *state, command command) error {
	users, err := state.db.GetUsers(context.Background())

	if err != nil {
		return fmt.Errorf("faild to query all users %w", err)
	}

	for _, user := range users {
		if state.config.CurrentUserName == user.Name {
			fmt.Printf(" * %s (current)\n", user.Name)
		} else {
			fmt.Printf(" * %s\n", user.Name)
		}
	}

	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID: %s\n", user.ID)
	fmt.Printf(" * Name: %s\n", user.Name)
}
