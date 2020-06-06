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

// DeploymentAdd defines the command for creating a deployment.
var DeploymentAdd = &cli.Command{
	Name:        "deployment",
	Description: "Use this command to add a deployment.",
	Usage:       "Add a new deployment from the provided configuration",
	Action:      deploymentAdd,
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

		&cli.StringFlag{
			EnvVars: []string{"VELA_REF", "DEPLOYMENT_REF"},
			Name:    "ref",
			Usage:   "provide the reference to deploy - this can be a branch, commit (SHA) or tag",
			Value:   "refs/heads/master",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TARGET", "DEPLOYMENT_TARGET"},
			Name:    "target",
			Aliases: []string{"t"},
			Usage:   "provide the name for the target deployment environment",
			Value:   "production",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_DESCRIPTION", "DEPLOYMENT_DESCRIPTION"},
			Name:    "description",
			Aliases: []string{"d"},
			Usage:   "provide the description for the deployment",
			Value:   "Deployment request from Vela",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TASK", "DEPLOYMENT_TASK"},
			Name:    "task",
			Aliases: []string{"tk"},
			Usage:   "Provide the task for the deployment",
			Value:   "deploy:vela",
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
  1. Add a deployment for a repository.
    $ {{.HelpName}} --org github --repo octocat
  2. Add a deployment with a specific target environment.
    $ {{.HelpName}} --org github --repo octocat --target stage
  3. Add a deployment with a specific branch reference.
    $ {{.HelpName}} --org github --repo octocat --ref dev
  4. Add a deployment with a specific commit reference.
    $ {{.HelpName}} --org github --repo octocat --ref 48afb5bdc41ad69bf22588491333f7cf71135163
  5. Add a deployment with a specific description.
    $ {{.HelpName}} --org github --repo octocat --description 'my custom message'
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// create a deployment.
func deploymentAdd(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the deployment configuration
	d := &deployment.Config{
		Action:      addAction,
		Org:         c.String("org"),
		Repo:        c.String("repo"),
		Description: c.String("description"),
		Ref:         c.String("ref"),
		Target:      c.String("target"),
		Task:        c.String("task"),
		Output:      c.String("output"),
	}

	// validate deployment configuration
	err = d.Validate()
	if err != nil {
		return err
	}

	// execute the add call for the deployment configuration
	return d.Add(client)
}
