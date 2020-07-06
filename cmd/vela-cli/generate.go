// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/action"

	"github.com/urfave/cli/v2"
)

// generateCmds defines the commands for producing resources.
var generateCmds = &cli.Command{
	Name:        "generate",
	Category:    "Resource Management",
	Aliases:     []string{"gn"},
	Description: "Use this command to generate resources for Vela.",
	Usage:       "Generate resources for Vela via subcommands",
	Subcommands: []*cli.Command{
		// add the sub command for producing a config file
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#ConfigGenerate
		action.ConfigGenerate,

		// add the sub command for producing a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#PipelineGenerate
		action.PipelineGenerate,
	},
}
