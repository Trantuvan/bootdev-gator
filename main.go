package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/trantuvan/bootdev-gator/internal/config"
	"github.com/trantuvan/bootdev-gator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {
	cfg, err := config.Read()

	if err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)

	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	defer db.Close()
	commands := commands{}
	state := state{config: &cfg, db: database.New(db)}
	commandLineArgs := os.Args

	if len(commandLineArgs) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	commandArgs := commandLineArgs[1:]
	command := command{name: commandArgs[0], args: commandArgs[1:]}
	commands.register("register", handlerRegister)
	commands.register("login", handlerLogin)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", middlewareLoggedIn(handlerBrowse))
	errRun := commands.run(&state, command)

	if errRun != nil {
		log.Fatalf("failed to run: %v", errRun)
	}
}
