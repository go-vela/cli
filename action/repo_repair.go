// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/repo"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// RepoRepair defines the command for repairing settings of a repository.
var RepoRepair = &cli.Command{
	Name:        "chown",
	Description: "Use this command to repair a damaged repository.",
	Usage:       "Repair settings of the provided repository",
	Action:      repoRepair,
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
  1. Repair a damaged repository.
    $ {{.HelpName}} --org github --repo octocat
  2. Repair a damaged repository with json output.
    $ {{.HelpName}} --org github --repo octocat --output json
  3. Repair a damaged repository when org and repo config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/repo/repair/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// repair settings of a repository.
func repoRepair(c *cli.Context) error {
	// parse the Vela client from the context
	client, err := client.Parse(c)
	if err != nil {
		return err
	}

	// create the repo configuration
	r := &repo.Config{
		Action: repairAction,
		Org:    c.String("org"),
		Name:   c.String("repo"),
		Output: c.String("output"),
	}

	// validate repo configuration
	err = r.Validate()
	if err != nil {
		return err
	}

	// execute the repair call for the repo configuration
	return r.Repair(client)
}
