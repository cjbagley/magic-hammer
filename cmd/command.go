package cmd

type Command interface {
	Init([]string) error
	Run() error
	SubCommand() string
}
