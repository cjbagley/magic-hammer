package cmd

func (cmd *ImageCommand) Argurments() []string {
	var args = []string{
		"-path",
		"~/.",
		"-filter",
		"Triangle",
		"-define",
		"filter:support=2",
		"-thumbnail",
		"70%",
		"-unsharp",
		"0.25x0.25+8+0.065",
		"-dither",
		"None",
		"-posterize",
		"136",
		"-quality",
		"82",
		"-interlace",
		"none",
		"-colorspace",
		"sRGB",
		"-strip",
		"-format",
		"webp",
		"-define",
		"webp:method=6",
		"how-to-be-productive.png",
	}

	return args
}
