// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"github.com/go-vela/cli/cmd/config"
	"github.com/go-vela/cli/cmd/repo"
	"github.com/go-vela/cli/cmd/secret"

	"github.com/urfave/cli"
)

// removeCmds defines the commands for deleting resources.
var removeCmds = cli.Command{
	Name:        "remove",
	Category:    "Resource Management",
	Aliases:     []string{"rm"},
	Description: "Use this command to remove resources for Vela.",
	Usage:       "Remove resources for Vela via subcommands",
	Subcommands: []cli.Command{
		config.RemoveCmd,
		repo.RemoveCmd,
		secret.RemoveCmd,
	},
}
