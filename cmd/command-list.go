package cmd

func CommandList() []Command {
	return []Command{
		NewImageCommand(),
		NewVideoCommand(),
	}
}
