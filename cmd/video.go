package cmd

import (
	"errors"
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
)

type VideoCommand struct {
	fs            *flag.FlagSet
	crfValue      int
	inputFilename string
	outputName    string
	fromSeconds   int
	fromMinutes   int
	toSeconds     int
	toMinutes     int
}

func (cmd *VideoCommand) SubCommand() string {
	return cmd.fs.Name()
}

func (cmd *VideoCommand) Init(args []string) error {
	return cmd.fs.Parse(args)
}

// Uses ffmpeg two pass process to convert video to webm format
func (cmd *VideoCommand) Run() error {
	if err := cmd.ValidateFlags(); err != nil {
		return err
	}

	// First Pass
	log.Println("Executing pass one....")
	passOne := exec.Command("ffmpeg", cmd.FirstPassArguments()...)
	passOne.Stdout = os.Stdout
	if err := passOne.Run(); err != nil {
		return err
	}
	log.Println("Finished pass one....")

	log.Println("Executing pass two....")
	passTwo := exec.Command("ffmpeg", cmd.SecondPassArguments()...)
	passTwo.Stdout = os.Stdout
	if err := passTwo.Run(); err != nil {
		return err
	}
	log.Println("Finished pass two....")

	// Clean up for 2 pass log
	delLog := exec.Command("rm", "ffmpeg*.log")
	if err := delLog.Run(); err != nil {
		return err
	}

	log.Println("Video converted!")

	return nil
}

func NewVideoCommand() Command {
	cmd := &VideoCommand{
		fs: flag.NewFlagSet("video", flag.ContinueOnError),
	}
	cmd.fs.IntVar(&cmd.crfValue, "crf", 60, "The crf value to use. 0 to 63. The lower the number, the higher the quality (and filesize).")
	cmd.fs.StringVar(&cmd.inputFilename, "f", "input.mp4", "The input file to process.")
	cmd.fs.StringVar(&cmd.outputName, "o", "output", "The output filename to use, excluding the file extension (webm).")
	cmd.fs.IntVar(&cmd.fromSeconds, "fs", 0, "The number of seconds to start the video from. Use in conjunction with 'fm' to cut any content before the given minutes/seconds.")
	cmd.fs.IntVar(&cmd.fromMinutes, "fm", 0, "The number of minutes to start the video from. Use in conjunction with 'fs' to cut any content before the given minutes/seconds.")
	cmd.fs.IntVar(&cmd.toSeconds, "ts", 0, "The number of seconds to end the video at. Use in conjunction with 'tm' to cut any content after the given minutes/seconds.")
	cmd.fs.IntVar(&cmd.toMinutes, "tm", 0, "The number of minutes to end the video an. Use in conjunction with 'ts' to cut any content after the given minutes/seconds.")

	return cmd
}

func (cmd *VideoCommand) ValidateFlags() error {
	var err []string
	if cmd.crfValue < 0 || cmd.crfValue > 63 {
		err = append(err, "crf value must be between 0 and 63")
	}

	if cmd.inputFilename == "" {
		err = append(err, "the input filename must not be blank")
	}

	if cmd.outputName == "" {
		err = append(err, "the output name must not be blank")
	}

	if cmd.fromSeconds < 0 || cmd.fromSeconds > 59 {
		err = append(err, "from seconds must be between 0 and 59")
	}

	if cmd.fromMinutes < 0 || cmd.fromMinutes > 59 {
		err = append(err, "from minutes must be between 0 and 59")
	}

	if cmd.toSeconds < 0 || cmd.toSeconds > 59 {
		err = append(err, "to seconds must be between 0 and 59")
	}

	if cmd.toMinutes < 0 || cmd.toMinutes > 59 {
		err = append(err, "to minutes must be between 0 and 59")
	}

	if len(err) > 0 {
		return errors.New(strings.Join(err, ", "))
	}

	return nil
}
