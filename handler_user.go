package main

import "fmt"

func handlerLogin(state *state, command command) error {
	if len(command.args) == 0 {
		return fmt.Errorf("command: login expected 1 arg username")
	}

	state.config.SetUser(command.args[0])
	fmt.Printf("username: %s has been set", state.config.CurrentUserName)
	return nil
}
