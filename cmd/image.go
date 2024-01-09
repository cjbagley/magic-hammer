package cmd

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"

	h "github.com/cjbagley/magic-hammer/helpers"
)

type ImageCommand struct {
	fs               *flag.FlagSet
	inputFilename    string
	outputFilename   string
	quality          int
	thumbnailPercent int
}

func (cmd *ImageCommand) SubCommand() string {
	return cmd.fs.Name()
}

func (cmd *ImageCommand) Init(args []string) error {
	if err := cmd.fs.Parse(args); err != nil {
		return err
	}

	cmd.outputFilename = "output.webp"
	n := strings.LastIndexByte(cmd.inputFilename, '.')
	if n != -1 && cmd.inputFilename[:n] != "" {
		cmd.outputFilename = cmd.inputFilename[:n] + ".webp"
	}
	if cmd.inputFilename == cmd.outputFilename {
		cmd.outputFilename = "converted-" + cmd.outputFilename
	}

	return nil
}

func (cmd *ImageCommand) Run() error {
	if err := cmd.ValidateFlags(); err != nil {
		return err
	}

	log.Println("Starting image conversion - stand by....")
	convert := exec.Command("convert", cmd.Argurments()...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	convert.Stdout = &out
	convert.Stderr = &stderr
	if err := convert.Run(); err != nil {
		return fmt.Errorf("%v:%v", err, stderr.String())
	}
	log.Println("Image converted")
	return nil
}

func NewImageCommand() *ImageCommand {
	cmd := &ImageCommand{
		fs: flag.NewFlagSet("image", flag.ContinueOnError),
	}
	cmd.fs.IntVar(&cmd.thumbnailPercent, "tp", 70, "The thumbnail percentage value to use. If 0, will not set thumbnail.")
	cmd.fs.IntVar(&cmd.quality, "q", 82, "The image quality value to use.")
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

	if len(err) > 0 {
		return errors.New(strings.Join(err, ", "))
	}

	return nil
}
