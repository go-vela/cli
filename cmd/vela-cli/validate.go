// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/command/pipeline"
)

// validateCmds defines the commands for validating resources.
var validateCmds = &cli.Command{
	Name:                   "validate",
	Category:               "Pipeline Management",
	Aliases:                []string{"vd"},
	Description:            "Use this command to validate a resource for Vela.",
	Usage:                  "Validate a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for validating a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/pipeline?tab=doc#CommandValidate
		pipeline.CommandValidate,
	},
}
