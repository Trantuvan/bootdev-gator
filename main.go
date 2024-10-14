package main

import (
	"log"
	"os"

	"github.com/trantuvan/bootdev-gator/internal/config"
)

type state struct {
	config *config.Config
}

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
