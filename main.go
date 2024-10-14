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
	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	commands := commands{}
	state := state{config: &cfg}
	commandLineArgs := os.Args

	if len(commandLineArgs) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	commandArgs := commandLineArgs[1:]
	loginCommand := command{name: commandArgs[0], args: commandArgs[1:]}
	commands.register(loginCommand.name, handlerLogin)
	errRun := commands.run(&state, loginCommand)

	if errRun != nil {
		log.Fatalf("failed to run: %v", errRun)
	}
}
