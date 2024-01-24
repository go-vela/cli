// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"fmt"
	"strconv"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/worker"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// CommandUpdate defines the command for modifying a worker.
var CommandUpdate = &cli.Command{
	Name:        "worker",
	Description: "(Platform Admin Only) Use this command to update a worker.",
	Usage:       "Update a worker from the provided configuration",
	Action:      update,
	Flags: []cli.Flag{

		// Worker Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_WORKER_ACTIVE", "WORKER_ACTIVE"},
			Name:    internal.FlagActive,
			Aliases: []string{"a"},
			Usage:   "current status of the worker",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_WORKER_ADDRESS", "WORKER_ADDRESS"},
			Name:    internal.FlagWorkerAddress,
			Aliases: []string{"wa"},
			Usage:   "provide the address of the worker as a fully qualified url (<scheme>://<host>)",
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_WORKER_BUILD_LIMIT", "WORKER_BUILD_LIMIT"},
			Name:    "build-limit",
			Aliases: []string{"bl"},
			Usage:   "build limit for the worker",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_WORKER_HOSTNAME", "WORKER_HOSTNAME"},
			Name:    internal.FlagWorkerHostname,
			Aliases: []string{"wh"},
			Usage:   "provide the hostname of the worker",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_WORKER_ROUTES", "WORKER_ROUTES"},
			Name:    "routes",
			Aliases: []string{"r"},
			Usage:   "route assignment for the worker",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "WORKER_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Update a worker to change its build limit to 2.
    $ {{.HelpName}} --worker.hostname MyWorker --build-limit 2
  2. Update a worker with custom routes.
    $ {{.HelpName}} --worker.hostname MyWorker --route large --route small

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/worker/update/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to modify a worker.
func update(c *cli.Context) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// parse the Vela client from the context
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/client?tab=doc#Parse
	client, err := client.Parse(c)
	if err != nil {
		return err
	}

	var active bool

	if len(c.String(internal.FlagActive)) > 0 {
		active, err = strconv.ParseBool(c.String(internal.FlagActive))
		if err != nil {
			return err
		}
	}

	// create the worker configuration
	w := &worker.Config{
		Hostname:   c.String(internal.FlagWorkerHostname),
		Address:    c.String(internal.FlagWorkerAddress),
		Active:     &active,
		Routes:     c.StringSlice("routes"),
		BuildLimit: c.Int64("build-limit"),
	}

	// validate worker configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/worker?tab=doc#Config.Validate
	err = w.Validate()
	if err != nil {
		return err
	}

	// execute the update call for the worker configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/worker?tab=doc#Config.Update
	return w.Update(client)
}
