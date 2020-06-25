// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/deployment"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// DeploymentView defines the command for inspecting a deployment.
var DeploymentView = &cli.Command{
	Name:        "deployment",
	Description: "Use this command to view a deployment.",
	Usage:       "View details of the provided deployment",
	Action:      deploymentView,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "DEPLOYMENT_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "provide the organization for the deployment",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "DEPLOYMENT_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "provide the repository for the deployment",
		},

		// Deployment Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_DEPLOYMENT", "DEPLOYMENT_NUMBER"},
			Name:    "deployment",
			Aliases: []string{"d"},
			Usage:   "provide the number for the deployment",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "DEPLOYMENT_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "print the output in default, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View deployment details for a repository.
    $ {{.HelpName}} --org github --repo octocat --deployment 1
  2. View deployment details for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --deployment 1 --output json
  3. View deployment details for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --deployment 1

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/deployment/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// inspect a deployment.
func deploymentView(c *cli.Context) error {
	// parse the Vela client from the context
	client, err := client.Parse(c)
	if err != nil {
		return err
	}

	// create the deployment configuration
	d := &deployment.Config{
		Action: viewAction,
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Number: c.Int("deployment"),
		Output: c.String("output"),
	}

	// validate deployment configuration
	err = d.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the deployment configuration
	return d.View(client)
}
