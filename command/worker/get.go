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

	"github.com/urfave/cli/v2"
)

// CommandGet defines the command for capturing a list of workers.
var CommandGet = &cli.Command{
	Name:        "worker",
	Description: "Use this command to get a list of workers.",
	Usage:       "Display a list workers",
	Action:      get,
	Flags: []cli.Flag{

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
  1. Get a list of workers.
    $ {{.HelpName}}
  2. Get a list of workers with wide view output.
    $ {{.HelpName}} --output wide
  3. Get a list of workers with yaml output.
    $ {{.HelpName}} --output yaml
  4. Get a list of workers with json output.
    $ {{.HelpName}} --output json
  5. Get a list of workers when config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/worker/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to capture a list
// of workers.
func get(c *cli.Context) error {
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
		Action: internal.ActionGet,
		Output: c.String(internal.FlagOutput),
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
	return w.Get(client)
}
