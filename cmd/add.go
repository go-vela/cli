// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"github.com/go-vela/cli/cmd/repo"
	"github.com/go-vela/cli/cmd/secret"

	"github.com/urfave/cli/v2"
)

// addCmds defines the commands for adding resources.
var addCmds = cli.Command{
	Name:        "add",
	Category:    "Resource Management",
	Aliases:     []string{"a"},
	Description: "Use this command to add resources for Vela.",
	Usage:       "Add resources for Vela via subcommands",
	Subcommands: []*cli.Command{
		&repo.AddCmd,
		&secret.AddCmd,
	},
}
