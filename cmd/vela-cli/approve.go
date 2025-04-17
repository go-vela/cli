// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/command/build"
)

// approveCmds defines the commands for approving resources.
var approveCmds = &cli.Command{
	Name:                   "approve",
	Category:               "Resource Management",
	Description:            "Use this command to approve a resource for Vela.",
	Usage:                  "Approve a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Commands: []*cli.Command{
		// add the sub command for approving a build
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/build?tab=doc#CommandApprove
		build.CommandApprove,
	},
}
