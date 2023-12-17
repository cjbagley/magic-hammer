package cmd

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type VideoCommand struct {
	fs       *flag.FlagSet
	crfValue int
}

func (cmd *VideoCommand) SubCommand() string {
	return cmd.fs.Name()
}

func (cmd *VideoCommand) Init(args []string) error {
	return cmd.fs.Parse(args)
}

func (cmd *VideoCommand) ValidateFlags() error {
	var err []string
	if cmd.crfValue < 0 || cmd.crfValue > 63 {
		err = append(err, "crf value must be between 0 and 63")
	}

	if len(err) > 0 {
		return errors.New(strings.Join(err, ", "))
	}

	return nil
}

func (cmd *VideoCommand) Run() error {
	if err := cmd.ValidateFlags(); err != nil {
		return err
	}

	command := `ffmpeg -i input.mp4 -c:v libvpx-vp9 -b:v 0 -crf 60 -ss 00:00:14 -t 00:00:27 -map_metadata -1 -pass 1 -an -f null /dev/null && \
    ffmpeg -i input.mp4 -c:v libvpx-vp9 -b:v 0 -crf 60 -ss 00:00:14 -t 00:00:27 -map_metadata -1 -pass 2 -c:a libvorbis output.webm`
	fmt.Printf(command)
	return nil
}

func NewVideoCommand() Command {
	cmd := &VideoCommand{
		fs: flag.NewFlagSet("video", flag.ContinueOnError),
	}
	cmd.fs.IntVar(&cmd.crfValue, "crf", 60, "The crf value to use. 0 to 63. The lower the number, the higher the quality (and filesize).")

	return cmd
}
