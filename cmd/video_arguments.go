package cmd

import (
	"fmt"
	"strconv"
)

func (cmd *VideoCommand) FirstPassArguments() []string {
	var args = []string{
		"-i",
		cmd.inputFilename,
		"-c:v",
		"libvpx-vp9",
		"-b:v",
		"0",
		"-crf",
		strconv.Itoa(cmd.crfValue),
	}

	if toBeSet, time := cmd.GetStartTime(); toBeSet != false {
		args = append(args, "-ss", time)
	}
	if toBeSet, time := cmd.GetFinishTime(); toBeSet != false {
		args = append(args, "-t", time)
	}

	args = append(args,
		"-map_metadata",
		"-1",
		"-pass",
		"1",
		"-an",
		"-f",
		"null",
		// First pass needs to be set to null, not the output file
		"/dev/null",
	)

	return args
}

func (cmd *VideoCommand) SecondPassArguments() []string {
	var args = []string{
		"-i",
		cmd.inputFilename,
		"-c:v",
		"libvpx-vp9",
		"-b:v",
		"0",
		"-crf",
		strconv.Itoa(cmd.crfValue),
	}

	if toBeSet, time := cmd.GetStartTime(); toBeSet != false {
		args = append(args, "-ss", time)
	}
	if toBeSet, time := cmd.GetFinishTime(); toBeSet != false {
		args = append(args, "-t", time)
	}

	// End with the output file
	args = append(args,
		"-map_metadata",
		"-1",
		"-pass",
		"2",
		"-c:a",
		"libvorbis",
		cmd.outputName+".webm",
	)

	return args
}

func (cmd *VideoCommand) GetStartTime() (toBeSet bool, time string) {
	if cmd.fromSeconds == 0 && cmd.fromMinutes == 0 {
		return false, ""
	}

	return true, cmd.FormatTimeArgsToString(cmd.fromMinutes, cmd.fromSeconds)
}

func (cmd *VideoCommand) GetFinishTime() (toBeSet bool, time string) {
	if cmd.toSeconds == 0 && cmd.toMinutes == 0 {
		return false, ""
	}

	return true, cmd.FormatTimeArgsToString(cmd.toMinutes, cmd.toSeconds)
}

func (cmd *VideoCommand) FormatTimeArgsToString(minutes int, seconds int) string {
	return fmt.Sprintf("00:%02d:%02d", minutes, seconds)
}
