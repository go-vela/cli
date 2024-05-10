// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code with chown and repair
package schedule

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/schedule"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandRemove defines the command for removing a schedule.
var CommandRemove = &cli.Command{
	Name:        "schedule",
	Description: "Use this command to remove a schedule.",
	Usage:       "Remove the provided schedule",
	Action:      remove,
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

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "SCHEDULE_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Remove a schedule from a repository.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --schedule daily
  2. Remove a schedule from a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --schedule daily --output json
  3. Remove a schedule from a repository when config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/schedule/remove/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input and create the object used to remove a repository.
func remove(c *cli.Context) error {
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
		Action: internal.ActionRemove,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Name:   c.String(internal.FlagSchedule),
		Output: c.String(internal.FlagOutput),
		Color:  output.ColorOptionsFromCLIContext(c),
	}

	// validate schedule configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/schedule?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the remove call for the schedule configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/schedule?tab=doc#Config.Remove
	return s.Remove(client)
}
