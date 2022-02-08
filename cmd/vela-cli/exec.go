// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/command/pipeline"
	"github.com/go-vela/cli/internal"

	"github.com/urfave/cli/v2"
)

// execCmds defines the commands for executing resources.
var execCmds = &cli.Command{
	Name:                   internal.ActionExec,
	Category:               "Pipeline Management",
	Description:            "Use this command to execute a resource for Vela.",
	Usage:                  "Execute a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for executing a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/pipeline?tab=doc#CommandExec
		pipeline.CommandExec,
	},
}
