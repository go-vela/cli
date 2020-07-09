// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/action"

	"github.com/urfave/cli/v2"
)

// getCmds defines the commands for getting a list of resources.
var getCmds = &cli.Command{
	Name:                   "get",
	Category:               "Resource Management",
	Aliases:                []string{"g"},
	Description:            "Use this command to get a list of resources for Vela.",
	Usage:                  "Get a list of resources for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for getting a list of builds
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#BuildGet
		action.BuildGet,

		// add the sub command for getting a list of deployments
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#DeploymentGet
		action.DeploymentGet,

		// add the sub command for getting a list of hooks
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#HookGet
		action.HookGet,

		// add the sub command for getting a list of build logs
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#LogGet
		action.LogGet,

		// add the sub command for getting a list of repositories
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#RepoGet
		action.RepoGet,

		// add the sub command for getting a list of secrets
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#SecretGet
		action.SecretGet,

		// add the sub command for getting a list of services
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#ServiceGet
		action.ServiceGet,

		// add the sub command for getting a list of steps
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action?tab=doc#StepGet
		action.StepGet,
	},
}
