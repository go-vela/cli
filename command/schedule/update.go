// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code with add
package schedule

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/schedule"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/urfave/cli/v2"
)

// CommandUpdate defines the command for modifying a schedule.
var CommandUpdate = &cli.Command{
	Name:        "schedule",
	Description: "Use this command to update a schedule.",
	Usage:       "Update a schedule from the provided configuration",
	Action:      update,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "SCHEDULE_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the schedule",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "SCHEDULE_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the schedule",
		},

		// Schedule Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_SCHEDULE", "SCHEDULE_NAME"},
			Name:    internal.FlagSchedule,
			Aliases: []string{"s"},
			Usage:   "provide the name for the schedule",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_ACTIVE", "SCHEDULE_ACTIVE"},
			Name:    "active",
			Aliases: []string{"a"},
			Usage:   "provided the status for the schedule",
			Value:   "true",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_ENTRY", "SCHEDULE_ENTRY"},
			Name:    "entry",
			Aliases: []string{"e"},
			Usage:   "provide the entry for the schedule",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_BRANCH", "SCHEDULE_BRANCH"},
			Name:    "branch",
			Aliases: []string{"b"},
			Usage:   "provide the branch for the schedule",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "SCHEDULE_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Update a schedule for a repository with active disabled.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --schedule hourly --active false
  2. Update a schedule for a repository with a new entry.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --schedule nightly --entry '@nightly'
  3. Update a schedule for a repository when config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/schedule/update/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input and create the object used to modify a schedule.
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

	// create the schedule configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/schedule?tab=doc#Config
	s := &schedule.Config{
		Action: internal.ActionUpdate,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Active: c.Bool("active"),
		Name:   c.String(internal.FlagSchedule),
		Entry:  c.String("entry"),
		Output: c.String(internal.FlagOutput),
		Branch: c.String(internal.FlagBranch),
	}

	// validate schedule configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/schedule?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the update call for the schedule configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/schedule?tab=doc#Config.Update
	return s.Update(client)
}
