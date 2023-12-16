package cmd

import (
	"flag"
	"fmt"
)

type VideoCommand struct {
	fs   *flag.FlagSet
	test string
}

func (cmd *VideoCommand) SubCommand() string {
	return cmd.fs.Name()
}

func (cmd *VideoCommand) Init(args []string) error {
	return cmd.fs.Parse(args)
}

func (cmd *VideoCommand) Run() error {
	fmt.Printf(cmd.test)
	return nil
}

func NewVideoCommand() Command {
	cmd := &VideoCommand{
		fs: flag.NewFlagSet("video", flag.ContinueOnError),
	}
	cmd.fs.StringVar(&cmd.test, "test", "test", "test")

	return cmd
}
