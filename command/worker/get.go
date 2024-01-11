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

// CommandGet defines the command for capturing a list of workers.
var CommandGet = &cli.Command{
	Name:        "worker",
	Aliases:     []string{"workers"},
	Description: "Use this command to get a list of workers.",
	Usage:       "Display a list of workers",
	Action:      get,
	Flags: []cli.Flag{
		// Filter Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ACTIVE", "WORKER_ACTIVE"},
			Name:    internal.FlagActive,
			Aliases: []string{"a"},
			Usage:   "return workers based on active status",
		},

		// Time Flags

		&cli.Int64Flag{
			EnvVars: []string{"VELA_CHECKED_IN_BEFORE", "CHECKED_IN_BEFORE"},
			Name:    internal.FlagBefore,
			Aliases: []string{"bf", "until"},
			Usage:   "before time constraint",
		},
		&cli.Int64Flag{
			EnvVars: []string{"VELA_CHECKED_IN_AFTER", "CHECKED_IN_AFTER"},
			Name:    internal.FlagAfter,
			Aliases: []string{"af", "since"},
			Usage:   "after time constraint",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "WORKER_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in wide, json, spew or yaml",
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

	var active bool

	if len(c.String(internal.FlagActive)) > 0 {
		active, err = strconv.ParseBool(c.String(internal.FlagActive))
		if err != nil {
			return err
		}
	}

	// create the worker configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/worker?tab=doc#Config
	w := &worker.Config{
		Action:          internal.ActionGet,
		Active:          &active,
		Output:          c.String(internal.FlagOutput),
		CheckedInBefore: c.Int64(internal.FlagBefore),
		CheckedInAfter:  c.Int64(internal.FlagAfter),
	}

	// validate worker configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/worker?tab=doc#Config.Validate
	err = w.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the worker configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/worker?tab=doc#Config.Get
	return w.Get(client)
}
