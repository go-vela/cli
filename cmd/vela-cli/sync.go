// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/command/repo"

	"github.com/urfave/cli/v2"
)

// syncCmds defines the commands for syncing resources.
var syncCmds = &cli.Command{
	Name:                   "sync",
	Category:               "Resource Management",
	Aliases:                []string{"s"},
	Description:            "Use this command to sync Vela Database with SCM",
	Usage:                  "Sync database and SCM for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for sync repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/config?tab=doc#CommandSync
		repo.CommandSync,
	},
}
