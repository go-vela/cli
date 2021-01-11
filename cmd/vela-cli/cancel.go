// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/action"

	"github.com/urfave/cli/v2"
)

// cancelCmds defines the commands for canceling resources.
var cancelCmds = &cli.Command{
	Name:                   "cancel",
	Category:               "Pipeline Management",
	Aliases:                []string{"cx"},
	Description:            "Use this command to cancel a resource for Vela.",
	Usage:                  "Cancel a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for canceling a build
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#BuildCancel
		action.BuildCancel,
	},
}
