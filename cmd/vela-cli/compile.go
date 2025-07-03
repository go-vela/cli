// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/command/pipeline"
)

// compileCmds defines the commands for compiling resources.
var compileCmds = &cli.Command{
	Name:                   "compile",
	Category:               "Pipeline Management",
	Description:            "Use this command to compile a resource for Vela.",
	Usage:                  "Compile a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Commands: []*cli.Command{
		// add the sub command for compiling a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/pipeline?tab=doc#CommandCompile
		pipeline.CommandCompile,
	},
}
