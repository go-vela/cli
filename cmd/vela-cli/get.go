// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/command/build"
	"github.com/go-vela/cli/command/dashboard"
	"github.com/go-vela/cli/command/deployment"
	"github.com/go-vela/cli/command/hook"
	"github.com/go-vela/cli/command/log"
	"github.com/go-vela/cli/command/pipeline"
	"github.com/go-vela/cli/command/repo"
	"github.com/go-vela/cli/command/schedule"
	"github.com/go-vela/cli/command/secret"
	"github.com/go-vela/cli/command/service"
	"github.com/go-vela/cli/command/step"
	"github.com/go-vela/cli/command/worker"
)

// getCmds defines the commands for getting a list of resources.
var getCmds = &cli.Command{
	Name:                   "get",
	Category:               "Resource Management",
	Aliases:                []string{"g"},
	Description:            "Use this command to get a list of resources for Vela.",
	Usage:                  "Get a list of resources for Vela via subcommands",
	UseShortOptionHandling: true,
	Commands: []*cli.Command{
		// add the sub command for getting a list of builds
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/build?tab=doc#CommandGet
		build.CommandGet,

		// add the sub command for getting a list of user dashboards
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/dashboard?tab=doc#CommandGet
		dashboard.CommandGet,

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

		// add the sub command for getting a list of pipelines
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/pipeline?tab=doc#CommandGet
		pipeline.CommandGet,

		// add the sub command for getting a list of repositories
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/repo?tab=doc#CommandGet
		repo.CommandGet,

		// add the sub command for getting a list of schedules
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/schedule?tab=doc#CommandGet
		schedule.CommandGet,

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

		// add the sub command for getting a list of workers
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/worker?tab=doc#CommandGet
		worker.CommandGet,
	},
}
