package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

// Build time values.
//
// We will set this value via `go build -ldflags "-X main.Version"`
var (
	Version string
)

var app = cli.App{
	Name:        "byctl",
	Description: "the command-line tool for all storage services",
	Version:     Version,
	Flags:       mergeFlags(globalFlags),
	Commands: []*cli.Command{
		cpCmd,
		lsCmd,
		profileCmd,
		rmCmd,
		statCmd,
		teeCmd,
		catCmd,
		mvCmd,
		signCmd,
	},
}

func main() {
	err := app.Run(os.Args)
	if err != nil {
		// FIXME: we need to respect platform style later.
		os.Exit(1)
	}
}

func userConfigDir() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("$HOME is not specified")
	}
	return configDir
}
