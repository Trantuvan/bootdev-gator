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

func handlerReset(state *state, command command) error {
	err := state.db.DeleteUsers(context.Background())

	if err != nil {
		return err
	}

	fmt.Println("reset users tbl successful!")
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID: %s\n", user.ID)
	fmt.Printf(" * Name: %s\n", user.Name)
}
