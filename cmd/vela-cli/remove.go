// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/command/config"
	"github.com/go-vela/cli/command/repo"
	"github.com/go-vela/cli/command/secret"

	"github.com/urfave/cli/v2"
)

// removeCmds defines the commands for deleting resources.
var removeCmds = &cli.Command{
	Name:                   "remove",
	Category:               "Resource Management",
	Aliases:                []string{"r"},
	Description:            "Use this command to remove a resource for Vela.",
	Usage:                  "Remove a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for remove a config file
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/config?tab=doc#CommandRemove
		config.CommandRemove,

		// add the sub command for remove a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/repo?tab=doc#CommandRemove
		repo.CommandRemove,

		// add the sub command for remove a secret
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/secret?tab=doc#CommandRemove
		secret.CommandRemove,
	},
}
