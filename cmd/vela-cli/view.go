// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package main

import (
	"github.com/go-vela/cli/command/build"
	"github.com/go-vela/cli/command/config"
	"github.com/go-vela/cli/command/deployment"
	"github.com/go-vela/cli/command/hook"
	"github.com/go-vela/cli/command/log"
	"github.com/go-vela/cli/command/pipeline"
	"github.com/go-vela/cli/command/repo"
	"github.com/go-vela/cli/command/secret"
	"github.com/go-vela/cli/command/service"
	"github.com/go-vela/cli/command/step"
	"github.com/go-vela/cli/command/worker"

	"github.com/urfave/cli/v2"
)

// viewCmds defines the commands for inspecting resources.
var viewCmds = &cli.Command{
	Name:                   "view",
	Category:               "Resource Management",
	Aliases:                []string{"v"},
	Description:            "Use this command to view a resource for Vela.",
	Usage:                  "View details for a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Subcommands: []*cli.Command{
		// add the sub command for getting a list of builds
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/build?tab=doc#CommandView
		build.CommandView,

		// add the sub command for inspecting a config file
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/config?tab=doc#CommandView
		config.CommandView,

		// add the sub command for getting a list of deployments
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/deployment?tab=doc#CommandView
		deployment.CommandView,

		// add the sub command for getting a list of hooks
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/build?tab=doc#CommandView
		hook.CommandView,

		// add the sub command for getting a list of build logs
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/log?tab=doc#CommandView
		log.CommandView,

		// add the sub command for inspecting a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/pipeline?tab=doc#CommandView
		pipeline.CommandView,

		// add the sub command for getting a list of repositories
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/repo?tab=doc#CommandView
		repo.CommandView,

		// add the sub command for getting a list of secrets
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/secret?tab=doc#CommandView
		secret.CommandView,

		// add the sub command for getting a list of services
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/service?tab=doc#CommandView
		service.CommandView,

		// add the sub command for getting a list of steps
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/step?tab=doc#CommandView
		step.CommandView,

		// add the sub command for viewing a worker
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/worker?tab=doc#CommandView
		worker.CommandView,
	},
}
