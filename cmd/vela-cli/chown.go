// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/command/repo"

	"github.com/urfave/cli/v2"
)

// chownCmds defines the commands for changing ownership of a resource.
var chownCmds = &cli.Command{
	Name:                   "chown",
	Category:               "Repository Management",
	Aliases:                []string{"c"},
	Description:            "Use this command to change ownership of a resource for Vela.",
	Usage:                  "Change ownership of resources for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for changing ownership of a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/repo?tab=doc#CommandChown
		repo.CommandChown,
	},
}
