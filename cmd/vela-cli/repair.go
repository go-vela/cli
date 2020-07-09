// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/action"

	"github.com/urfave/cli/v2"
)

// repairCmds defines the commands for repairing resources.
var repairCmds = &cli.Command{
	Name:        "repair",
	Category:    "Resource Management",
	Aliases:     []string{"rp"},
	Description: "Use this command to repair a resource for Vela.",
	Usage:       "Repair a resource for Vela via subcommands",
	Subcommands: []*cli.Command{
		// add the sub command for repairing a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#RepoRepair
		action.RepoRepair,
	},
}
