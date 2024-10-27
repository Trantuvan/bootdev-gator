package main

import (
	"context"
	"fmt"
)

func handlerReset(state *state, command command) error {
	err := state.db.DeleteUsers(context.Background())

	if err != nil {
		return fmt.Errorf("couldn't delete users %w", err)
	}

	fmt.Println("reset users tbl successful!")
	return nil
}
