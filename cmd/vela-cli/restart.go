// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/command/build"
)

// restartCmds defines the commands for restarting resources.
var restartCmds = &cli.Command{
	Name:                   "restart",
	Category:               "Resource Management",
	Aliases:                []string{"rs"},
	Description:            "Use this command to restart a resource for Vela.",
	Usage:                  "Restart a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for restarting a build
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/build?tab=doc#CommandRestart
		build.CommandRestart,
	},
}
