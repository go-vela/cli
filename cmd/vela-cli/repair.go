// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/command/repo"
)

// repairCmds defines the commands for repairing resources.
var repairCmds = &cli.Command{
	Name:                   "repair",
	Category:               "Repository Management",
	Aliases:                []string{"rp"},
	Description:            "Use this command to repair a resource for Vela.",
	Usage:                  "Repair a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Commands: []*cli.Command{
		// add the sub command for repairing a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/repo?tab=doc#CommandRepair
		repo.CommandRepair,
	},
}
