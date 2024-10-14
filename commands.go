package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands map[string]func(*state, command) error

func (commands commands) register(name string, f func(*state, command) error) {
	commands[name] = f
}

func (commands commands) run(state *state, command command) error {
	if state == nil {
		return fmt.Errorf("command: %s state is not exiest", command.name)
	}

	cmd, ok := commands[command.name]

	if !ok {
		return fmt.Errorf("command: %s is not exiest", command.name)
	}

	if err := cmd(state, command); err != nil {
		return err
	}

	return nil
}
