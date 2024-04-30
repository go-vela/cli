// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/command/pipeline"
)

// expandCmds defines the commands for expanding resources.
var expandCmds = &cli.Command{
	Name:                   "expand",
	Category:               "Pipeline Management",
	Description:            "Use this command to expand a resource for Vela.",
	Usage:                  "Expand a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for expanding a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/pipeline?tab=doc#CommandExpand
		pipeline.CommandExpand,
	},
}
