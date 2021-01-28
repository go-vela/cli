// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/action"

	"github.com/urfave/cli/v2"
)

// addCmds defines the commands for creating resources.
var addCmds = &cli.Command{
	Name:                   "add",
	Category:               "Resource Management",
	Aliases:                []string{"a"},
	Description:            "Use this command to add resources to Vela.",
	Usage:                  "Add resources to Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for creating a deployment
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#DeploymentAdd
		action.DeploymentAdd,

		// add the sub command for creating a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#DeploymentAdd
		action.RepoAdd,

		// add the sub command for creating a secret
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#DeploymentAdd
		action.SecretAdd,
	},
}
