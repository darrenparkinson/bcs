package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	// log.SetOutput(ioutil.Discard)
	c := cli.NewCLI("bcs-cli", "0.0.1")
	c.Args = os.Args[1:]
	for _, arg := range c.Args {
		if arg == "-v" || arg == "--version" {
			newArgs := make([]string, len(c.Args)+1)
			newArgs[0] = "version"
			copy(newArgs[1:], c.Args)
			c.Args = newArgs
			break
		}
	}

	c.Commands = Commands

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(exitStatus)
}
