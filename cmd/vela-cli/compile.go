// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/go-vela/cli/command/pipeline"

	"github.com/urfave/cli/v2"
)

// compileCmds defines the commands for compiling resources.
var compileCmds = &cli.Command{
	Name:                   "compile",
	Category:               "Pipeline Management",
	Description:            "Use this command to compile a resource for Vela.",
	Usage:                  "Compile a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for compiling a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/pipeline?tab=doc#CommandCompile
		pipeline.CommandCompile,
	},
}
