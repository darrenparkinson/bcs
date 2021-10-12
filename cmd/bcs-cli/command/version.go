package command

import (
	"bytes"
	"fmt"

	"github.com/mitchellh/cli"
)

// VersionCommand is the top level struct for the cli VersionCommand.
// It holds a reference to the cli.Ui for logging etc.
type VersionCommand struct {
	Revision          string
	Version           string
	VersionPrerelease string
	Ui                cli.Ui
}

// Help provies the help text for this command.
func (c *VersionCommand) Help() string {
	return ""
}

// Run provides the command functionality
func (c *VersionCommand) Run(_ []string) int {
	var versionString bytes.Buffer
	fmt.Fprintf(&versionString, "bcs-cli v%s", c.Version)
	if c.VersionPrerelease != "" {
		fmt.Fprintf(&versionString, ".%s", c.VersionPrerelease)

		if c.Revision != "" {
			fmt.Fprintf(&versionString, " (%s)", c.Revision)
		}
	}

	c.Ui.Output(versionString.String())
	return 0
}

// Synopsis provides the one liner
func (c *VersionCommand) Synopsis() string {
	return "Show version information."
}
