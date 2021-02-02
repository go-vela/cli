// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// nolint: dupl // ignore similar code among actions
package action

import (
	"fmt"

	"github.com/go-vela/cli/action/repo"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// RepoRepair defines the command for repairing settings of a repository.
var RepoRepair = &cli.Command{
	Name:        "repo",
	Description: "Use this command to repair a damaged repository.",
	Usage:       "Repair settings of the provided repository",
	Action:      repoRepair,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "REPO_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the repository",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "REPO_NAME"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the name for the repository",
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
  1. Repair a damaged repository.
    $ {{.HelpName}} --org MyOrg --repo MyRepo
  2. Repair a damaged repository with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --output json
  3. Repair a damaged repository when config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/repo/repair/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// repair settings of a repository.
func repoRepair(c *cli.Context) error {
	// load variables from the config file
	err := load(c)
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

	// create the repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config
	r := &repo.Config{
		Action: repairAction,
		Org:    c.String(internal.FlagOrg),
		Name:   c.String(internal.FlagRepo),
		Output: c.String(internal.FlagOutput),
	}

	// validate repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Validate
	err = r.Validate()
	if err != nil {
		return err
	}

	// execute the repair call for the repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Repair
	return r.Repair(client)
}
