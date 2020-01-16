// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"github.com/go-vela/cli/cmd/config"
	"github.com/go-vela/cli/cmd/pipe"

	"github.com/urfave/cli"
)

// genCmds defines the commands for adding resources.
var genCmds = cli.Command{
	Name:        "generate",
	Category:    "Resource Management",
	Aliases:     []string{"gen"},
	Description: "Use this command to generate local files",
	Usage:       "Generate resources for Vela via subcommands",
	Subcommands: []cli.Command{
		config.GenCmd,
		pipe.GenCmd,
	},
}
