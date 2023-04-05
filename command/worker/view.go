// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

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

// CommandView defines the command for inspecting a worker.
var CommandView = &cli.Command{
	Name:        "worker",
	Description: "(Admin Only) Use this command to view a worker.",
	Usage:       "View details of the provided worker",
	Action:      view,
	Flags: []cli.Flag{

		// Worker Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_WORKER_HOSTNAME", "WORKER_HOSTNAME"},
			Name:    internal.FlagWorkerHostname,
			Aliases: []string{"wh"},
			Usage:   "provide the hostname of the worker",
			// there is no current enforcement on passing a valid hostname
			// if none is supplied, use a random id.
			Value: uuid.NewString(),
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_WORKER_REGISTRATION_TOKEN", "WORKER_REGISTRATION_TOKEN"},
			Name:    internal.FlagWorkerRegistrationToken,
			Aliases: []string{"wr"},
			Usage:   "toggle to show the registration token for the worker",
			Value:   "false",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "WORKER_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View details of a worker.
    $ {{.HelpName}} --worker.hostname MyWorker
  2. View registration token for a worker.
    $ {{.HelpName}} --worker.registration.token true
  3. View details of a worker with json output.
    $ {{.HelpName}} --worker.hostname MyWorker --output json
  4. View details of a worker when config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/worker/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to inspect a worker.
//
//nolint:dupl // ignore similar code with chown, get, remove and repair
func view(c *cli.Context) error {
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
		Action:            internal.ActionView,
		Hostname:          c.String(internal.FlagWorkerHostname),
		RegistrationToken: c.Bool(internal.FlagWorkerRegistrationToken),
		Output:            c.String(internal.FlagOutput),
	}

	// validate worker configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/worker?tab=doc#Config.Validate
	err = w.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the worker configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/worker?tab=doc#Config.View
	return w.View(client)
}
