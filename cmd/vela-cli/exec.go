// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/command/pipeline"
)

// execCmds defines the commands for executing resources.
var execCmds = &cli.Command{
	Name:                   "exec",
	Category:               "Pipeline Management",
	Description:            "Use this command to execute a resource for Vela.",
	Usage:                  "Execute a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Commands: []*cli.Command{
		// add the sub command for executing a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/pipeline?tab=doc#CommandExec
		pipeline.CommandExec,
	},
}
