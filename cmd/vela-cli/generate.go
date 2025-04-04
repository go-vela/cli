// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/command/completion"
	"github.com/go-vela/cli/command/config"
	"github.com/go-vela/cli/command/docs"
	"github.com/go-vela/cli/command/pipeline"
)

// generateCmds defines the commands for producing resources.
var generateCmds = &cli.Command{
	Name:                   "generate",
	Category:               "Resource Management",
	Aliases:                []string{"gn"},
	Description:            "Use this command to generate resources for Vela.",
	Usage:                  "Generate resources for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for producing a shell auto completion script
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/completion?tab=doc#CommandGenerate
		completion.CommandGenerate,

		// add the sub command for producing a config file
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/config?tab=doc#CommandGenerate
		config.CommandGenerate,

		// add the sub command for producing documentation
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/docs?tab=doc#CommandGenerate
		docs.CommandGenerate,

		// add the sub command for producing a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/pipeline?tab=doc#CommandGenerate
		pipeline.CommandGenerate,
	},
}
