// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/command/dashboard"
	"github.com/go-vela/cli/command/deployment"
	"github.com/go-vela/cli/command/repo"
	"github.com/go-vela/cli/command/schedule"
	"github.com/go-vela/cli/command/secret"
	"github.com/go-vela/cli/command/worker"
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
		// add the sub command for creating a dashboard
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/dashboard?tab=doc#CommandAdd
		dashboard.CommandAdd,
		// add the sub command for creating a deployment
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/deployment?tab=doc#CommandAdd
		deployment.CommandAdd,

		// add the sub command for creating a repository
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/repo?tab=doc#CommandAdd
		repo.CommandAdd,

		// add the sub command for creating a schedule
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/schedule?tab=doc#CommandAdd
		schedule.CommandAdd,

		// add the sub command for creating a secret
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/secret?tab=doc#CommandAdd
		secret.CommandAdd,

		// add the sub command for creating a worker
		//
		// https://pkg.go.dev/github.com/go-vela/cli/command/worker?tab=doc#CommandAdd
		worker.CommandAdd,
	},
}
