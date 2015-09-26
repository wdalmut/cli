package cli

import (
	"bytes"
	"fmt"
)

type HelpFunc func(map[string]CommandFactory) string

func BasicHelpFunc(app string) HelpFunc {
	return func(factories map[string]CommandFactory) string {
		var buffer bytes.Buffer

		buffer.WriteString(fmt.Sprintf("usage: %s [--version] [--help] <command> [<args>]", app))

		if factories != nil {
			buffer.WriteString("\n\n")

			for name, factory := range factories {
				command := factory()
				buffer.WriteString(fmt.Sprintf("%s %s\n", name, command.Synopsis()))
			}
		}

		return buffer.String()
	}
}
