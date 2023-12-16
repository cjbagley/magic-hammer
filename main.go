package main

import (
	"errors"
	"fmt"
	"os"

	commands "github.com/cjbagley/magic-hammer/cmd"
	"github.com/cjbagley/magic-hammer/helpers"
)

func run(args []string) error {
	if len(args) < 1 {
		return errors.New("Please pass a command")
	}

	subcommand := os.Args[1]

	for _, cmd := range commands.CommandList() {
		if cmd.SubCommand() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknown command %s", subcommand)
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		helpers.Exit(err.Error())
	}
}
