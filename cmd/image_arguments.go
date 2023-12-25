package cmd

import "strconv"

func (cmd *ImageCommand) Argurments() []string {
	var args = []string{
		"-path",
		".",
		"-filter",
		"Triangle",
		"-define",
		"filter:support=2",
		"-unsharp",
		"0.25x0.25+8+0.065",
		"-dither",
		"None",
		"-posterize",
		"136",
		"-interlace",
		"none",
		"-colorspace",
		"sRGB",
		"-strip",
		"-format",
		"webp",
		"-define",
		"webp:method=6",
		"-quality",
		strconv.Itoa(cmd.quality),
	}

	if toBeSet, percentage := cmd.GetThumbnailPercent(); toBeSet == true {
		args = append(args, "-thumbnail", percentage)
	}

	args = append(args, "-write", cmd.outputName+".webp")

	// Final arguments - the file to process
	args = append(args, cmd.inputFilename)

	return args
}

func (cmd *ImageCommand) GetThumbnailPercent() (toBeSet bool, percent string) {
	if cmd.thumbnailPercent == 0 {
		return false, ""
	}

	return true, strconv.Itoa(cmd.thumbnailPercent) + "%"
}
