// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/command/build"
	"github.com/go-vela/cli/command/deployment"
	"github.com/go-vela/cli/command/hook"
	"github.com/go-vela/cli/command/log"
	"github.com/go-vela/cli/command/repo"
	"github.com/go-vela/cli/command/secret"
	"github.com/go-vela/cli/command/service"
	"github.com/go-vela/cli/command/step"
	"github.com/go-vela/cli/internal"

	"github.com/urfave/cli/v2"
)

// getCmds defines the commands for getting a list of resources.
var getCmds = &cli.Command{
	Name:                   internal.ActionGet,
	Category:               "Resource Management",
	Aliases:                []string{"g"},
	Description:            "Use this command to get a list of resources for Vela.",
	Usage:                  "Get a list of resources for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for getting a list of builds
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/build?tab=doc#CommandGet
		build.CommandGet,

		// add the sub command for getting a list of deployments
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/deployment?tab=doc#CommandGet
		deployment.CommandGet,

		// add the sub command for getting a list of hooks
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/build?tab=doc#CommandGet
		hook.CommandGet,

		// add the sub command for getting a list of build logs
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/log?tab=doc#CommandGet
		log.CommandGet,

		// add the sub command for getting a list of repositories
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/repo?tab=doc#CommandGet
		repo.CommandGet,

		// add the sub command for getting a list of secrets
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/secret?tab=doc#CommandGet
		secret.CommandGet,

		// add the sub command for getting a list of services
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/service?tab=doc#CommandGet
		service.CommandGet,

		// add the sub command for getting a list of steps
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/step?tab=doc#CommandGet
		step.CommandGet,
	},
}
