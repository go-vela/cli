// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/command/pipeline"

	"github.com/urfave/cli/v2"
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
