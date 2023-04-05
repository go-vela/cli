// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

//nolint:dupl // ignore similar code with update
package worker

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/worker"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
)

// CommandAdd defines the command for adding a worker.
var CommandAdd = &cli.Command{
	Name:        "worker",
	Description: "(Admin Only) Use this command to add a worker.",
	Usage:       "Add a new worker from the provided configuration",
	Action:      add,
	Flags: []cli.Flag{

		// Worker Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_WORKER_ADDRESS", "WORKER_ADDRESS"},
			Name:    internal.FlagWorkerAddress,
			Aliases: []string{"wa"},
			Usage:   "provide the address of the worker",
		},

		&cli.StringFlag{
			EnvVars: []string{"VELA_WORKER_HOSTNAME", "WORKER_HOSTNAME"},
			Name:    internal.FlagWorkerHostname,
			Aliases: []string{"wh"},
			Usage:   "provide the hostname of the worker (defaults is random ID)",
			// there is no current enforcement on passing a valid hostname
			// if none is supplied, use a random id.
			Value: uuid.NewString(),
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
  1. Add a worker reachable at the give address.
    $ {{.HelpName}} --worker.address myworker.example.com

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/worker/add/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to create a worker.
func add(c *cli.Context) error {
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

	// create the worker configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/worker?tab=doc#Config
	w := &worker.Config{
		Action:   internal.ActionAdd,
		Address:  c.String(internal.FlagWorkerAddress),
		Hostname: c.String(internal.FlagWorkerHostname),
		Output:   c.String(internal.FlagOutput),
	}

	// validate worker configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/worker?tab=doc#Config.Validate
	err = w.Validate()
	if err != nil {
		return err
	}

	// execute the add call for the worker configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/worker?tab=doc#Config.Add
	return w.Add(client)
}
