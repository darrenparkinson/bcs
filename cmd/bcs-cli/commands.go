package main

import (
	"os"

	"github.com/darrenparkinson/bcs/cmd/bcs-cli/command"
	"github.com/mitchellh/cli"
)

// Commands holds a map of each of the cli commands
var Commands map[string]cli.CommandFactory

func init() {
	bui := &cli.BasicUi{Writer: os.Stdout}
	ui := &cli.ColoredUi{
		Ui:          bui,
		OutputColor: cli.UiColorBlue,
		InfoColor:   cli.UiColorGreen,
		WarnColor:   cli.UiColorYellow,
		ErrorColor:  cli.UiColorRed,
	}
	Commands = map[string]cli.CommandFactory{
		"download": func() (cli.Command, error) {
			return &command.DownloadCommand{Ui: ui}, nil
		},
		"parse": func() (cli.Command, error) {
			return &command.ParseFileCommand{Ui: ui}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Revision:          GitCommit,
				Version:           Version,
				VersionPrerelease: VersionPrerelease,
				Ui:                ui,
			}, nil
		},
	}
}
