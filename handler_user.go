package main

import "fmt"

func handlerLogin(state *state, command command) error {
	if len(command.args) == 0 {
		return fmt.Errorf("command: login expected 1 arg username")
	}

	if err := state.config.SetUser(command.args[0]); err != nil {
		return fmt.Errorf("failed to set user: %w", err)
	}

	fmt.Printf("username: %s has been set", state.config.CurrentUserName)
	return nil
}
