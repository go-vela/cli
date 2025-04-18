// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/command/config"
	"github.com/go-vela/cli/command/dashboard"
	"github.com/go-vela/cli/command/repo"
	"github.com/go-vela/cli/command/schedule"
	"github.com/go-vela/cli/command/secret"
	"github.com/go-vela/cli/command/settings"
	"github.com/go-vela/cli/command/user"
	"github.com/go-vela/cli/command/worker"
)

// updateCmds defines the commands for modifying resources.
var updateCmds = &cli.Command{
	Name:                   "update",
	Category:               "Resource Management",
	Aliases:                []string{"u"},
	Description:            "Use this command to update a resource for Vela.",
	Usage:                  "Update a resource for Vela via subcommands",
	UseShortOptionHandling: true,
	Commands: []*cli.Command{
		// add the sub command for modifying a config file
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/config?tab=doc#CommandUpdate
		config.CommandUpdate,

		// add the sub command for modifying a dashboard
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/dashboard?tab=doc#CommandUpdate
		dashboard.CommandUpdate,

		// add the sub command for modifying a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/repo?tab=doc#CommandUpdate
		repo.CommandUpdate,

		// add the sub command for modifying a schedule
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/schedule?tab=doc#CommandUpdate
		schedule.CommandUpdate,

		// add the sub command for modifying a secret
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/secret?tab=doc#CommandUpdate
		secret.CommandUpdate,

		// add the sub command for modifying settings
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/settings?tab=doc#CommandUpdate
		settings.CommandUpdate,

		// add the sub command for modifying a user
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/user?tab=doc#CommandUpdate
		user.CommandUpdate,

		// add the sub command for modifying a worker
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/worker?tab=doc#CommandUpdate
		worker.CommandUpdate,
	},
}
