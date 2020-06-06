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

// RepoRemove defines the command for removing a repository.
var RepoRemove = &cli.Command{
	Name:        "repo",
	Description: "Use this command to remove a repository.",
	Usage:       "Remove the provided repository",
	Action:      repoRemove,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "REPO_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "provide the organization for the repository",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "REPO_NAME"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "provide the name for the repository",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "REPO_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "print the output in default, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Remove a repository.
    $ {{.HelpName}} --org github --repo octocat
 2. Remove a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --output json
 3. Remove a repository when org and repo config or environment variables are set.
    $ {{.HelpName}}
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// remove a repository.
func repoRemove(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the repo configuration
	r := &repo.Config{
		Action: removeAction,
		Org:    c.String("org"),
		Name:   c.String("repo"),
		Output: c.String("output"),
	}

	// validate repo configuration
	err = r.Validate()
	if err != nil {
		return err
	}

	// execute the remove call for the repo configuration
	return r.Remove(client)
}
