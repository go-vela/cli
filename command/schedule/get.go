// SPDX-License-Identifier: Apache-2.0

package schedule

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/schedule"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandGet defines the command for capturing a list of schedules.
var CommandGet = &cli.Command{
	Name:        "schedule",
	Aliases:     []string{"schedules"},
	Description: "Use this command to get a list of schedules.",
	Usage:       "Display a list of schedules",
	Action:      get,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ORG", "SCHEDULE_ORG"),
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the schedule",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REPO", "SCHEDULE_REPO"),
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the schedule",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "SCHEDULE_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},

		// Pagination Flags

		&cli.IntFlag{
			Sources: cli.EnvVars("VELA_PAGE", "SCHEDULE_PAGE"),
			Name:    internal.FlagPage,
			Aliases: []string{"p"},
			Usage:   "print a specific page of schedules",
			Value:   1,
		},
		&cli.IntFlag{
			Sources: cli.EnvVars("VELA_PER_PAGE", "SCHEDULE_PER_PAGE"),
			Name:    internal.FlagPerPage,
			Aliases: []string{"pp"},
			Usage:   "number of schedules to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get a list of schedules for a repository.
    $ {{.FullName}} --org MyOrg --repo MyRepo
  2. Get a list of schedules for a repository with wide view output.
    $ {{.FullName}} --org MyOrg --repo MyRepo --output wide
  3. Get a list of schedules for a repository with yaml output.
    $ {{.FullName}} --org MyOrg --repo MyRepo --output yaml
  4. Get a list of schedules for a repository with json output.
    $ {{.FullName}} --org MyOrg --repo MyRepo --output json
  5. Get a list of schedules for a repository when config or environment variables are set.
    $ {{.FullName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/schedule/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input and create the object used to capture a list of schedules.
func get(_ context.Context, c *cli.Command) error {
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

	// create the schedule configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/schedule?tab=doc#Config
	s := &schedule.Config{
		Action:  internal.ActionGet,
		Org:     c.String(internal.FlagOrg),
		Repo:    c.String(internal.FlagRepo),
		Page:    c.Int(internal.FlagPage),
		PerPage: c.Int(internal.FlagPerPage),
		Output:  c.String(internal.FlagOutput),
		Color:   output.ColorOptionsFromCLIContext(c),
	}

	// validate schedule configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/schedule?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the schedule configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/schedule?tab=doc#Config.Get
	return s.Get(client)
}
