// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/action"

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
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#ConfigRemove
		action.ConfigRemove,

		// add the sub command for remove a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#RepoRemove
		action.RepoRemove,

		// add the sub command for remove a secret
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#SecretRemove
		action.SecretRemove,
	},
}
