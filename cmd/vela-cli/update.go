// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/action"

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
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#ConfigUpdate
		action.ConfigUpdate,

		// add the sub command for modifying a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#RepoUpdate
		action.RepoUpdate,

		// add the sub command for modifying a secret
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#SecretUpdate
		action.SecretUpdate,
	},
}
