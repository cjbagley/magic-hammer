package cmd

import (
	"flag"
)

type ImageCommand struct {
	fs               *flag.FlagSet
	thumbnailPercent int
	quality          int
}

func (cmd *ImageCommand) SubCommand() string {
	return cmd.fs.Name()
}

func (cmd *ImageCommand) Init(args []string) error {
	return cmd.fs.Parse(args)
}

func (cmd *ImageCommand) Run() error {
	return nil
}

func NewImageCommand() *ImageCommand {
	cmd := &ImageCommand{
		fs: flag.NewFlagSet("image", flag.ContinueOnError),
	}
	cmd.fs.IntVar(&cmd.thumbnailPercent, "tp", 70, "The thumbnail percentage value to use. If 0, will not set thumbnail.")
	cmd.fs.IntVar(&cmd.quality, "q", 82, "The image quality value to use.")

	return cmd
}
