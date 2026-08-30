package main

import (
	"fmt"
	"os"

	"github.com/dynamicmines/tic-tac-toe-arena/internal/cli"
	"github.com/dynamicmines/tic-tac-toe-arena/internal/game"
)

func main() {
	cfg, err := cli.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		cli.PrintUsage()
		os.Exit(1)
	}

	game.Run(cfg)
}