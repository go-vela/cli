// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/repo"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/constants"

	"github.com/urfave/cli/v2"
)

// RepoUpdate defines the command for modifying a repository.
var RepoUpdate = &cli.Command{
	Name:        "repo",
	Description: "Use this command to update a repository.",
	Usage:       "Update a new repository from the provided configuration",
	Action:      repoUpdate,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG"},
			Name:    "org",
			Usage:   "Provide the organization for the repository",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO"},
			Name:    "repo",
			Usage:   "Provide the name for the repository",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_BRANCH"},
			Name:    "branch",
			Usage:   "default branch for the repository",
			Value:   "master",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_LINK"},
			Name:    "link",
			Usage:   "full URL to repository in source control",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_CLONE"},
			Name:    "clone",
			Usage:   "full clone URL to repository in source control",
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_VISIBILITY"},
			Name:    "timeout",
			Usage:   "fine tune the access level required to view the repo",
			Value:   30,
		},
		&cli.Int64Flag{
			EnvVars: []string{"VELA_TIMEOUT"},
			Name:    "timeout",
			Usage:   "max time allowed per build in repository",
			Value:   30,
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_PRIVATE"},
			Name:    "timeout",
			Usage:   "max time allowed per build in repository",
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_PRIVATE"},
			Name:    "timeout",
			Usage:   "max time allowed per build in repository",
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_PRIVATE"},
			Name:    "timeout",
			Usage:   "max time allowed per build in repository",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_EVENT"},
			Name:    "event",
			Usage:   "webhook events repository responds to",
			Value: cli.NewStringSlice(
				constants.EventPush,
				constants.EventPull,
			),
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT"},
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Print the output in default, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Update a repository with push and pull request enabled.
   $ {{.HelpName}} --org github --repo octocat --event push --event pull_request
 2. Update a repository with all event types enabled.
   $ {{.HelpName}} --org github --repo octocat --event push --event pull_request --event tag --event deployment --event comment
 3. Update a repository with a longer build timeout.
   $ {{.HelpName}} --org github --repo octocat --timeout 90
 4. Update a repository with push and pull request enabled when org and repo config or environment variables are set.
   $ {{.HelpName}} --event push --event pull_request
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// modify a repository.
func repoUpdate(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the repo configuration
	r := &repo.Config{
		Action:     updateAction,
		Org:        c.String("org"),
		Name:       c.String("repo"),
		Branch:     c.String("branch"),
		Link:       c.String("link"),
		Clone:      c.String("clone"),
		Visibility: c.String("visibility"),
		Timeout:    c.Int64("timeout"),
		Private:    c.Bool("private"),
		Trusted:    c.Bool("trusted"),
		Active:     c.Bool("active"),
		Events:     c.StringSlice("event"),
		Output:     c.String("output"),
	}

	// validate repo configuration
	err = r.Validate()
	if err != nil {
		return err
	}

	// execute the update call for the repo configuration
	return r.Update(client)
}
