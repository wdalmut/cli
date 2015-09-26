package cli

type Command interface {
	Help() string
	Run(args []string) int
	Synopsis() string
}

type CommandFactory func() Command
