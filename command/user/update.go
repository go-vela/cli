// SPDX-License-Identifier: Apache-2.0

package user

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/user"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandUpdate defines the command for updating a user.
var CommandUpdate = &cli.Command{
	Name:        "user",
	Description: "Use this command to update a user.",
	Usage:       "Update a user from the provided configuration",
	Action:      update,
	Flags: []cli.Flag{

		// User Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_USER_NAME", "USER_NAME"},
			Name:    internal.FlagName,
			Usage:   "provide the name of the user",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_USER_ADD_FAVORITES", "USER_ADD_FAVORITES"},
			Name:    "add-favorites",
			Usage:   "provide the list of repositories to add as favorites for the user",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_USER_DROP_FAVORITES", "USER_DROP_FAVORITES"},
			Name:    "drop-favorites",
			Usage:   "provide the list of repositories to remove from favorites of the user",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_USER_ADD_DASHBOARDS", "USER_ADD_DASHBOARDS"},
			Name:    "add-dashboards",
			Usage:   "provide the list of UUIDs for dashboards to add to the user",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_USER_DROP_DASHBOARDS", "USER_DROP_DASHBOARDS"},
			Name:    "drop-dashboards",
			Usage:   "provide the list of UUIDs for dashboareds to remove from the user",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "REPO_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Update current user to add a repository to favorites.
    $ {{.HelpName}} --add-favorites Org-1/Repo-1
  2. Update current user to remove a repository from favorites.
    $ {{.HelpName}} --drop-favorites Org-1/Repo-1
  3. Update current user to add a dashboard.
    $ {{.HelpName}} --add-dashboards c8da1302-07d6-11ea-882f-4893bca275b8
  4. Update current user to remove a dashboard.
    $ {{.HelpName}} --drop-dashboards c8da1302-07d6-11ea-882f-4893bca275b8

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/user/update/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to update a user.
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

	// create the user configuration
	u := &user.Config{
		Name:           c.String(internal.FlagName),
		AddFavorites:   c.StringSlice("add-favorites"),
		DropFavorites:  c.StringSlice("drop-favorites"),
		AddDashboards:  c.StringSlice("add-dashboards"),
		DropDashboards: c.StringSlice("drop-dashboards"),
		Output:         c.String(internal.FlagOutput),
		Color:          output.ColorOptionsFromCLIContext(c),
	}

	// execute the update call for the user configuration
	return u.Update(client)
}
