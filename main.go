package main

import (
	"fmt"

	"github.com/trantuvan/bootdev-gator/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("trantuvan")
	cfg = config.Read()
	fmt.Printf("config %+v", cfg)
}
