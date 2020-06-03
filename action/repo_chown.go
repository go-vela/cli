// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/repo"

	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// RepoChown defines the command for changing ownership of a repository.
var RepoChown = &cli.Command{
	Name:        "chown",
	Description: "Use this command to change the repository owner.",
	Usage:       "Change ownership of the provided repository",
	Action:      repoChown,
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

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT"},
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Print the output in wide, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Change ownership of a repository.
    $ {{.HelpName}} --org github --repo octocat
 2. Change ownership of a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --output json
 3. Change ownership of a repository when org and repo config or environment variables are set.
    $ {{.HelpName}}
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// change ownership of a repository.
func repoChown(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the repo configuration
	r := &repo.Config{
		Action: chownAction,
		Org:    c.String("org"),
		Name:   c.String("repo"),
		Output: c.String("output"),
	}

	// validate repo configuration
	err = r.Validate()
	if err != nil {
		return err
	}

	// execute the chown call for the repo configuration
	return r.Chown(client)
}
