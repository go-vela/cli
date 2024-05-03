// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/command/build"
	"github.com/go-vela/cli/command/config"
	"github.com/go-vela/cli/command/deployment"
	"github.com/go-vela/cli/command/hook"
	"github.com/go-vela/cli/command/log"
	"github.com/go-vela/cli/command/pipeline"
	"github.com/go-vela/cli/command/repo"
	"github.com/go-vela/cli/command/schedule"
	"github.com/go-vela/cli/command/secret"
	"github.com/go-vela/cli/command/service"
	"github.com/go-vela/cli/command/settings"
	"github.com/go-vela/cli/command/step"
	"github.com/go-vela/cli/command/worker"
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
		// add the sub command for viewing a build
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/build?tab=doc#CommandView
		build.CommandView,

		// add the sub command for viewing a config file
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/config?tab=doc#CommandView
		config.CommandView,

		// add the sub command for viewing a deployment
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/deployment?tab=doc#CommandView
		deployment.CommandView,

		// add the sub command for viewing a hook
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/build?tab=doc#CommandView
		hook.CommandView,

		// add the sub command for viewing a build log
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/log?tab=doc#CommandView
		log.CommandView,

		// add the sub command for viewing a pipeline
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/pipeline?tab=doc#CommandView
		pipeline.CommandView,

		// add the sub command for viewing a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/repo?tab=doc#CommandView
		repo.CommandView,

		// add the sub command for viewing a schedule
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/schedule?tab=doc#CommandView
		schedule.CommandView,

		// add the sub command for viewing a secret
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/secret?tab=doc#CommandView
		secret.CommandView,

		// add the sub command for viewing a service
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/service?tab=doc#CommandView
		service.CommandView,

		// add the sub command for viewing settings
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/settings?tab=doc#CommandView
		settings.CommandView,

		// add the sub command for viewing a step
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/step?tab=doc#CommandView
		step.CommandView,

		// add the sub command for viewing a worker
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/worker?tab=doc#CommandView
		worker.CommandView,
	},
}
