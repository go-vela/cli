// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"fmt"

	"github.com/go-vela/go-vela/vela"

	"github.com/urfave/cli"
)

// RemoveCmd defines the command for deleting a repository.
var RemoveCmd = cli.Command{
	Name:        "repo",
	Description: "Use this command to remove a repository.",
	Usage:       "Remove a repository",
	Action:      remove,
	Before:      validate,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		cli.StringFlag{
			Name:   "org",
			Usage:  "Provide the organization for the repository",
			EnvVar: "REPO_ORG",
		},
		cli.StringFlag{
			Name:   "repo",
			Usage:  "Provide the repository contained with the organization",
			EnvVar: "REPO_NAME",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Remove a repository.
    $ {{.HelpName}} --org github --repo octocat
 2. Remove a repository when org and repo config or environment variables are set.
    $ {{.HelpName}}
`, cli.CommandHelpTemplate),
}

// helper function to execute a remove repo cli command
func remove(c *cli.Context) error {

	// get org and repo information from cmd flags
	org, repo := c.String("org"), c.String("repo")

	// create a carval client
	client, err := vela.NewClient(c.GlobalString("addr"), nil)
	if err != nil {
		return err
	}

	// set token from context
	client.Authentication.SetTokenAuth(c.GlobalString("token"))

	_, _, err = client.Repo.Remove(org, repo)
	if err != nil {
		return err
	}

	fmt.Printf("repo \"%s/%s\" was removed \n", org, repo)

	return nil
}
