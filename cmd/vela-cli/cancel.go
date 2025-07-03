// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/command/build"
)

// cancelCmds defines the commands for canceling resources.
var cancelCmds = &cli.Command{
	Name:                   "cancel",
	Category:               "Resource Management",
	Aliases:                []string{"cx"},
	Description:            "Use this command to cancel a resource for Vela.",
	Usage:                  "Cancel a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Commands: []*cli.Command{
		// add the sub command for canceling a build
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/build?tab=doc#CommandCancel
		build.CommandCancel,
	},
}
