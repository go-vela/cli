// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code with update
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

// CommandAdd defines the command for creating a schedule.
var CommandAdd = &cli.Command{
	Name:        "schedule",
	Description: "Use this command to add a schedule.",
	Usage:       "Add a new schedule from the provided configuration",
	Action:      add,
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

		// Schedule Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_SCHEDULE", "SCHEDULE_NAME"),
			Name:    internal.FlagSchedule,
			Aliases: []string{"s"},
			Usage:   "provide the name for the schedule",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ACTIVE", "SCHEDULE_ACTIVE"),
			Name:    "active",
			Aliases: []string{"a"},
			Usage:   "provided the status for the schedule",
			Value:   "true",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ENTRY", "SCHEDULE_ENTRY"),
			Name:    "entry",
			Aliases: []string{"e"},
			Usage:   "provide the entry for the schedule",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_BRANCH", "SCHEDULE_BRANCH"),
			Name:    "branch",
			Aliases: []string{"b"},
			Usage:   "provide the branch for the schedule",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "SCHEDULE_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Add a schedule to a repository with active not enabled.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --schedule hourly --entry '0 * * * *' --active false
  2. Add a schedule to a repository with a nightly entry.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --schedule nightly --entry '0 0 * * *'
  3. Add a schedule to a repository with a weekly entry.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --schedule weekly --entry '@weekly'
  4. Add a schedule to a repository when config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/schedule/add/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input and create the object used to create a schedule.
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

	// create the schedule configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/schedule?tab=doc#Config
	s := &schedule.Config{
		Action: internal.ActionAdd,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Active: c.Bool("active"),
		Name:   c.String(internal.FlagSchedule),
		Entry:  c.String("entry"),
		Output: c.String(internal.FlagOutput),
		Color:  output.ColorOptionsFromCLIContext(c),
		Branch: c.String(internal.FlagBranch),
	}

	// validate schedule configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/schedule?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the add call for the schedule configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/schedule?tab=doc#Config.Add
	return s.Add(client)
}
