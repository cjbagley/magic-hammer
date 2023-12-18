package cmd

import "strconv"

func (cmd *VideoCommand) FirstPassArguments() []string {
	return []string{
		"-i",
		cmd.inputFilename,
		"-c:v",
		"libvpx-vp9",
		"-b:v",
		"0",
		"-crf",
		strconv.Itoa(cmd.crfValue),
		// "-ss",
		// "00:00:14",
		//"-t",
		//"00:00:00",
		"-map_metadata",
		"-1",
		"-pass",
		"1",
		"-an",
		"-f",
		"null",
		"/dev/null",
	}
}

func (cmd *VideoCommand) SecondPassArguments() []string {
	return []string{
		"-i",
		cmd.inputFilename,
		"-c:v",
		"libvpx-vp9",
		"-b:v",
		"0",
		"-crf",
		strconv.Itoa(cmd.crfValue),
		// "-ss",
		// "00:00:14",
		//"-t",
		//"00:00:00",
		"-map_metadata",
		"-1",
		"-pass",
		"2",
		"-c:a",
		"libvorbis",
		cmd.outputName + ".webm",
	}
}
