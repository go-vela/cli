// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/command/build"

	"github.com/urfave/cli/v2"
)

// restartCmds defines the commands for restarting resources.
var restartCmds = &cli.Command{
	Name:                   "restart",
	Category:               "Resource Management",
	Aliases:                []string{"rs"},
	Description:            "Use this command to restart a resource for Vela.",
	Usage:                  "Restart a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for restarting a build
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/build?tab=doc#CommandRestart
		build.CommandRestart,
	},
}
