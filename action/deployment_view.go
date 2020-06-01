// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/deployment"

	"github.com/go-vela/sdk-go/vela"

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
			EnvVars: []string{"VELA_ORG"},
			Name:    "org",
			Usage:   "Provide the organization for the deployment",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO"},
			Name:    "repo",
			Usage:   "Provide the repository for the deployment",
		},

		// Deployment Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_DEPLOYMENT"},
			Name:    "deployment",
			Aliases: []string{"d"},
			Usage:   "Provide the deployment number",
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
 1. View deployment details for a repository.
    $ {{.HelpName}} --org github --repo octocat --deployment 1
 2. View deployment details for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --deployment 1 --output json
 3. View deployment details for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --deployment 1
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// inspect a deployment.
func deploymentView(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

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
