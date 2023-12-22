package cmd

import (
	"flag"
	"fmt"
)

type ImageCommand struct {
	fs   *flag.FlagSet
	name string
}

func (cmd *ImageCommand) SubCommand() string {
	return cmd.fs.Name()
}

func (cmd *ImageCommand) Init(args []string) error {
	return cmd.fs.Parse(args)
}

func (cmd *ImageCommand) Run() error {
	fmt.Printf(cmd.name)
	return nil
}

func NewImageCommand() *ImageCommand {
	cmd := &ImageCommand{
		fs: flag.NewFlagSet("image", flag.ContinueOnError),
	}
	cmd.fs.StringVar(&cmd.name, "name", "Colin", "name")

	return cmd
}
