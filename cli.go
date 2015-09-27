package cli

import (
	"io"
	"os"
)

type Cli struct {
	Name     string
	Version  string
	Args     []string
	Commands map[string]CommandFactory
	Help     HelpFunc
	Out      io.Writer
}

func NewCli(app, version string) *Cli {
	return &Cli{
		Name:    app,
		Version: version,
		Help:    BasicHelpFunc(app),
		Out:     os.Stdout,
	}
}

func (c *Cli) isVersion() bool {
	flag := Parse(c.Args)
	if flag.Bool("-v", "--version", "-version") {
		return true
	}

	return false
}

func (c *Cli) isHelp() bool {
	if len(c.Args) == 0 {
		return true
	}

	flag := Parse(c.Args)
	if flag.Bool("-h", "--help", "-help", "?") {
		return true
	}

	return false
}

func (c *Cli) Run() int {
	if c.isHelp() {
		for name, command := range c.Commands {
			for _, flag := range c.Args {
				if flag == name {
					c.Out.Write([]byte(command().Help()))
					return 0
				}
			}
		}
		c.Out.Write([]byte(c.Help(c.Commands) + "\n"))
		return 0
	}

	if c.isVersion() {
		c.Out.Write([]byte(c.Version + "\n"))
		return 0
	}

	if factory, ok := c.Commands[c.Args[0]]; ok {
		return factory().Run(c.Args[1:])
	}

	return 1
}
