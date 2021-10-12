package command

import (
	"context"
	"flag"
	"os"
	"strings"

	"github.com/darrenparkinson/bcs/pkg/ciscobcs"
	"github.com/mitchellh/cli"
)

// DownloadCommand is the top level struct for the cli DownloadCommand.
// It holds a reference to the cli.Ui for logging etc.
type DownloadCommand struct {
	Ui cli.Ui
}

// Help provies the help text for this command.
func (c *DownloadCommand) Help() string {
	helpText := `
Usage: bcs-cli [global options] download [options]

  Download bulk data file to local file.

  Downloads the bulk data for a given customer ID and API key. 

Options:
  -id=ID              The customer ID to download for. Required.

  -apikey=KEY            The API Key to use for the download. Required.

  -filename=FILENAME  Specify the filename to save the jsonlines
                      output to. Default bcs_bulk.jsonl.

`
	return strings.TrimSpace(helpText)
}

// Run provides the command functionality
func (c *DownloadCommand) Run(args []string) int {
	var customerID, apikey, filename string
	cmdFlags := flag.NewFlagSet("download", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Ui.Output(c.Help()) }
	cmdFlags.StringVar(&customerID, "id", "280987866", "customer id (default demo customer)")
	cmdFlags.StringVar(&apikey, "apikey", "lqWTZbApQZgSR52ag8NS9jc5STobf6hMjm3Kyf30", "api key with access to customer")
	cmdFlags.StringVar(&filename, "filename", "bcs_bulk.jsonl", "download file to create")
	//TODO: Add separate or raw option and maybe download separate files
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	bcs, err := ciscobcs.NewClient(apikey, nil)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	file, err := os.Create(filename)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	err = bcs.BulkService.Download(context.Background(), customerID, file)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	return 0
}

// Synopsis provides the one liner
func (c *DownloadCommand) Synopsis() string {
	return "Download raw data from the bulk endpoint and save as a file."
}
