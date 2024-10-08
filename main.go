package main

import (
	"fmt"
	"log"
	"os"

	"github.com/trantuvan/bootdev-gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands map[string]func(*state, command) error

func main() {
	cfg := config.Read()
	commands := commands{}
	state := state{config: &cfg}
	commandLineArgs := os.Args

	if len(commandLineArgs) < 2 {
		log.Fatal("you've must provide arg for command")
	}

	commandArgs := commandLineArgs[1:]
	loginCommand := command{name: commandArgs[0], args: commandArgs[1:]}
	commands.register(loginCommand.name, handlerLogin)
	err := commands.run(&state, loginCommand)

	if err != nil {
		log.Fatal(err)
	}
}

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

func handlerLogin(state *state, command command) error {
	if len(command.args) == 0 {
		return fmt.Errorf("command: login expected 1 arg username")
	}

	state.config.SetUser(command.args[0])
	fmt.Printf("username: %s has been set", state.config.CurrentUserName)
	return nil
}
