package command

import (
	"flag"
	"fmt"
	"strings"

	"github.com/darrenparkinson/bcs/pkg/ciscobcs"
	"github.com/mitchellh/cli"
)

// ParseFileCommand is the top level struct for the cli ParseFileCommand.
// It holds a reference to the cli.Ui for logging etc.
type ParseFileCommand struct {
	Ui cli.Ui
}

// Help provies the help text for this command.
func (c *ParseFileCommand) Help() string {
	helpText := `
Usage: bcs-cli [global options] parse [options]

  Parse bulk data file for stats.

  After downloading a bulk data file, you can use this command
  to parse the file and provide stats including the number of
  lines processed, a count of each type and a list of any
  unrecognised types. 

Options:
  -filename=FILENAME  Specify the filename for the jsonlines
                      file to process.

`
	return strings.TrimSpace(helpText)
}

// Run provides the command functionality
func (c *ParseFileCommand) Run(args []string) int {
	var filename string

	cmdFlags := flag.NewFlagSet("parse", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.StringVar(&filename, "filename", "bcs_bulk.jsonl", "filename of bulk download file to parse")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	results, err := ciscobcs.ParseBulkFile(filename)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	for _, e := range results.Errors {
		c.Ui.Error(fmt.Sprintf("error in results: %s", e))
	}
	if len(results.Errors) > 0 {
		c.Ui.Warn(fmt.Sprintf("%d errors above", len(results.Errors)))
	}
	c.Ui.Info(fmt.Sprintf("%d lines processed:", results.LineCount))
	for k, v := range results.CountOfTypes {
		c.Ui.Info(fmt.Sprintf("  * %s: %d", k, v))
	}
	for k, v := range results.UnrecognisedTypes {
		c.Ui.Warn(fmt.Sprintf("unrecognised type: %s: %d", k, v))
	}
	return 0
}

// Synopsis provides the one liner
func (c *ParseFileCommand) Synopsis() string {
	return "Parse downloaded bulk data file for stats."
}
