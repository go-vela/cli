// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"github.com/go-vela/cli/cmd/config"
	"github.com/go-vela/cli/cmd/repo"
	"github.com/go-vela/cli/cmd/secret"

	"github.com/urfave/cli/v2"
)

// updateCmds defines the command for updating resources.
var updateCmds = cli.Command{
	Name:        "update",
	Category:    "Resource Management",
	Aliases:     []string{"u"},
	Description: "Use this command to update resources for Vela.",
	Usage:       "Update resources for Vela via subcommands",
	Subcommands: []*cli.Command{
		&config.UpdateCmd,
		&repo.UpdateCmd,
		&secret.UpdateCmd,
	},
}
