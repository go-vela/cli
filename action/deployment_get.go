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

// DeploymentGet defines the command for capturing a list of deployments.
var DeploymentGet = &cli.Command{
	Name:        "deployment",
	Aliases:     []string{"deployments"},
	Description: "Use this command to get a list of deployments.",
	Usage:       "Display a list of deployments",
	Action:      deploymentGet,
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

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "DEPLOYMENT_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "print the output in default, wide, yaml or json format",
		},

		// Pagination Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_PAGE", "DEPLOYMENT_PAGE"},
			Name:    "page",
			Aliases: []string{"p"},
			Usage:   "print a specific page of deployments",
			Value:   1,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_PER_PAGE", "DEPLOYMENT_PER_PAGE"},
			Name:    "per.page",
			Aliases: []string{"pp"},
			Usage:   "number of deployments to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get deployments for a repository.
    $ {{.HelpName}} --org github --repo octocat
  2. Get deployments for a repository with wide view output.
    $ {{.HelpName}} --org github --repo octocat --output wide
  3. Get deployments for a repository with yaml output.
    $ {{.HelpName}} --org github --repo octocat --output yaml
  4. Get deployments for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --output json
  5. Get deployments for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/deployment/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// capture a list of deployments.
func deploymentGet(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the deployment configuration
	d := &deployment.Config{
		Action:  getAction,
		Org:     c.String("org"),
		Repo:    c.String("repo"),
		Page:    c.Int("page"),
		PerPage: c.Int("per.page"),
		Output:  c.String("output"),
	}

	// validate deployment configuration
	err = d.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the deployment configuration
	return d.Get(client)
}
