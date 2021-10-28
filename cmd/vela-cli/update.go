// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/command/config"
	"github.com/go-vela/cli/command/repo"
	"github.com/go-vela/cli/command/secret"

	"github.com/urfave/cli/v2"
)

// updateCmds defines the commands for modifying resources.
var updateCmds = &cli.Command{
	Name:                   "update",
	Category:               "Resource Management",
	Aliases:                []string{"u"},
	Description:            "Use this command to update a resource for Vela.",
	Usage:                  "Update a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for modifying a config file
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/config?tab=doc#CommandUpdate
		config.CommandUpdate,

		// add the sub command for modifying a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/repo?tab=doc#CommandUpdate
		repo.CommandUpdate,

		// add the sub command for modifying a secret
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/secret?tab=doc#CommandUpdate
		secret.CommandUpdate,
	},
}
