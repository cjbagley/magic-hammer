package cmd

import (
	"errors"
	"flag"
	"strings"

	h "github.com/cjbagley/magic-hammer/helpers"
)

type ImageCommand struct {
	fs               *flag.FlagSet
	inputFilename    string
	outputName       string
	quality          int
	thumbnailPercent int
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
	cmd.fs.StringVar(&cmd.outputName, "o", "output", "The output filename to use, excluding the file extension (webp).")
	cmd.fs.StringVar(&cmd.inputFilename, "f", "input.jpg", "The input file to process.")

	return cmd
}

func (cmd *ImageCommand) ValidateFlags() error {
	var err []string

	if !h.IsValidPercent(cmd.thumbnailPercent) {
		err = append(err, "thumbnail percent value must be between 0 and 100")
	}

	if !h.IsValidPercent(cmd.quality) {
		err = append(err, "quality must be between 0 and 100")
	}

	if !h.IsValidString(cmd.inputFilename) {
		err = append(err, "the input filename must not be blank")
	}

	if !h.IsValidString(cmd.outputName) {
		err = append(err, "the output name must not be blank")
	}

	if len(err) > 0 {
		return errors.New(strings.Join(err, ", "))
	}

	return nil
}
