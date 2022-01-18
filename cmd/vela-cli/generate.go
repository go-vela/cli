// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/command/completion"
	"github.com/go-vela/cli/command/config"
	"github.com/go-vela/cli/command/docs"
	"github.com/go-vela/cli/command/pipeline"

	"github.com/urfave/cli/v2"
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
