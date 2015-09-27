# Cli

A simple Cli package (merge of other CLI ideas)

```go
package main

import (
    "github.com/wdalmut/cli"
)

func main() {
    c := cli.NewCli("app", "1.0.0")
    c.Args = os.Args[1:]

	cli.Commands = map[string]CommandFactory{
		"foo": func() Command {
			return Command {
                Help: "This is a very long help message",
                Synopsis: "A short help message",
                Run: func(args []string) int {
                    flag := cli.Flag(args)
                    opt := flag.String("default value", "-f", "--force", "-force")

                    // continue

                    return 0 // exitcode
                },
            }
		},
    }

    exitStatus := c.Run()
    os.Exit(exitStatus)
}
```

