// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"github.com/go-vela/cli/cmd/build"

	"github.com/urfave/cli"
)

// restartCmds defines the command for restarting resources.
var restartCmds = cli.Command{
	Name:        "restart",
	Category:    "Pipeline Management",
	Aliases:     []string{"r"},
	Description: "Use this command to restart resources for Vela.",
	Usage:       "Restart build for a vela pipeline",
	Subcommands: []cli.Command{
		build.RestartCmd,
	},
}
