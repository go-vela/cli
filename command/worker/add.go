// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"context"
	"fmt"
	"net/url"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/worker"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandAdd defines the command for adding a worker.
var CommandAdd = &cli.Command{
	Name:        "worker",
	Description: "(Platform Admin Only) Use this command to add a worker.",
	Usage:       "Add a new worker from the provided configuration",
	Action:      add,
	Flags: []cli.Flag{

		// Worker Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_WORKER_ADDRESS", "WORKER_ADDRESS"),
			Name:    internal.FlagWorkerAddress,
			Aliases: []string{"wa"},
			Usage:   "provide the address of the worker as a fully qualified url (<scheme>://<host>)",
		},

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_WORKER_HOSTNAME", "WORKER_HOSTNAME"),
			Name:    internal.FlagWorkerHostname,
			Aliases: []string{"wh"},
			Usage:   "provide the hostname of the worker (defaults to hostname of worker address)",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "WORKER_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Add a worker reachable at the provided address.
    $ {{.FullName}} --worker.address https://myworker.example.com
  2. Add a worker reachable at the provided address with specific hostname.
    $ {{.FullName}} --worker.hostname MyWorker --worker.address https://myworker.example.com

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/worker/add/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to create a worker.
func add(ctx context.Context, c *cli.Command) error {
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
		Color:    output.ColorOptionsFromCLIContext(c),
	}

	// if no hostname was passed in, parse the hostname
	// out of the provided address
	if len(w.Hostname) == 0 {
		url, err := url.Parse(w.Address)
		if err != nil {
			return fmt.Errorf("address is not a valid url")
		}

		w.Hostname = url.Hostname()
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
	return w.Add(ctx, client)
}
