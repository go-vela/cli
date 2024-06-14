// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/dashboard"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandUpdate defines the command for updating a dashboard.
var CommandUpdate = &cli.Command{
	Name:        "dashboard",
	Description: "Use this command to update a dashboard.",
	Usage:       "Update a dashboard from the provided configuration",
	Action:      update,
	Flags: []cli.Flag{

		// Dashboard Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_DASHBOARD_ID", "DASHBOARD_ID"},
			Name:    "id",
			Usage:   "provide the uuid for the dashboard",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_DASHBOARD_NAME", "DASHBOARD_NAME"},
			Name:    "name",
			Usage:   "provide the name for the dashboard",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_DASHBOARD_ADD_REPOS", "DASHBOARD_ADD_REPOS"},
			Name:    "add-repos",
			Usage:   "provide the list of repositories to add for the dashboard",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_DASHBOARD_DROP_REPOS", "DASHBOARD_DROP_REPOS"},
			Name:    "drop-repos",
			Usage:   "provide the list of repositories to remove from the dashboard",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_DASHBOARD_TARGET_REPOS", "DASHBOARD_TARGET_REPOS"},
			Name:    "target-repos",
			Usage:   "provide the list of repositories to target for updates for the dashboard",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_DASHBOARD_ADD_ADMINS", "DASHBOARD_ADD_ADMINS"},
			Name:    "add-admins",
			Usage:   "provide the list of admins to add for the dashboard",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_DASHBOARD_DROP_ADMINS", "DASHBOARD_DROP_ADMINS"},
			Name:    "drop-admins",
			Usage:   "provide the list of admins to remove from the dashboard",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_DASHBOARD_REPOS_BRANCH", "DASHBOARD_REPOS_BRANCH"},
			Name:    "branches",
			Aliases: []string{"branch"},
			Usage:   "filter builds in all repositories by branch",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_DASHBOARD_REPOS_EVENT", "DASHBOARD_REPOS_EVENT"},
			Name:    "events",
			Aliases: []string{"event"},
			Usage:   "filter builds in all repositories by event",
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
  1. Update a dashboard to add a repository.
    $ {{.HelpName}} --id c8da1302-07d6-11ea-882f-4893bca275b8 --add-repos Org-1/Repo-1
  2. Update a dashboard to remove a repository.
    $ {{.HelpName}} --id c8da1302-07d6-11ea-882f-4893bca275b8 --drop-repos Org-1/Repo-1
  3. Update a dashboard to add event and branch filters to specific repositories.
    $ {{.HelpName}} --id c8da1302-07d6-11ea-882f-4893bca275b8 --target-repos Org-1/Repo-1,Org-2/Repo-2 --branches main --events push
  4. Update a dashboard to change the name.
    $ {{.HelpName}} --id c8da1302-07d6-11ea-882f-4893bca275b8 --name MyDashboard
  5. Update a dashboard to add an admin.
    $ {{.HelpName}} --id c8da1302-07d6-11ea-882f-4893bca275b8 --add-admins JohnDoe
  6. Update a dashboard to remove an admin.
    $ {{.HelpName}} --id c8da1302-07d6-11ea-882f-4893bca275b8 --drop-admins JohnDoe

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/dashboard/update/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to update a dashboard.
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

	// create the dashboard configuration
	d := &dashboard.Config{
		Action:      internal.ActionUpdate,
		ID:          c.String("id"),
		Name:        c.String("name"),
		AddRepos:    c.StringSlice("add-repos"),
		DropRepos:   c.StringSlice("drop-repos"),
		TargetRepos: c.StringSlice("target-repos"),
		Branches:    c.StringSlice("branches"),
		Events:      c.StringSlice("events"),
		AddAdmins:   c.StringSlice("add-admins"),
		DropAdmins:  c.StringSlice("drop-admins"),
		Output:      c.String(internal.FlagOutput),
		Color:       output.ColorOptionsFromCLIContext(c),
	}

	// validate dashboard configuration
	err = d.Validate()
	if err != nil {
		return err
	}

	// execute the update call for the dashboard configuration
	return d.Update(client)
}
