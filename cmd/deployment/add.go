// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"fmt"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/library"

	"github.com/urfave/cli/v2"
)

// AddCmd defines the command for adding a repository.
var AddCmd = cli.Command{
	Name:        "deployment",
	Description: "Use this command to add a deployment.",
	Usage:       "Add a deployment",
	Action:      add,
	Before:      validate,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		&cli.StringFlag{
			Name:    "org",
			Usage:   "Provide the organization for the repository",
			EnvVars: []string{"VELA_ORG"},
		},
		&cli.StringFlag{
			Name:    "repo",
			Usage:   "Provide the repository contained with the organization",
			EnvVars: []string{"VELA_REPO"},
		},
		&cli.StringFlag{
			Name:    "ref",
			Usage:   "Provide the reference to deploy - this can be a branch, commit (SHA) or tag",
			EnvVars: []string{"VELA_REF"},
			Value:   "master",
		},
		&cli.StringFlag{
			Name:    "target",
			Usage:   "Provide the name for the target deployment environment",
			EnvVars: []string{"VELA_TARGET"},
			Value:   "production",
		},

		// optional flags that can be supplied to a command
		&cli.StringFlag{
			Name:    "description",
			Usage:   "Provide the description for the deployment",
			EnvVars: []string{"VELA_DESCRIPTION"},
		},
		&cli.StringFlag{
			Name:    "task",
			Usage:   "Provide the task for the deployment",
			EnvVars: []string{"VELA_TASK"},
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Add a deployment for a repository.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld
 2. Add a deployment with a specific target environment.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --target stage
 3. Add a deployment with a specific branch reference.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --ref dev
 4. Add a deployment with a specific commit reference.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --ref 48afb5bdc41ad69bf22588491333f7cf71135163
 4. Add a deployment with a specific description.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --description 'my custom message'
`, cli.CommandHelpTemplate),
}

// helper function to execute a add repo cli command
func add(c *cli.Context) error {
	if len(c.String("ref")) == 0 {
		return util.InvalidCommand("ref")
	}

	if len(c.String("target")) == 0 {
		return util.InvalidCommand("target")
	}

	// get org, repo, ref and target information from cmd flags
	org, repo, ref, target := c.String("org"), c.String("repo"), c.String("ref"), c.String("target")

	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	client.Authentication.SetTokenAuth(c.String("token"))

	// resource to create on server
	request := &library.Deployment{
		Description: vela.String(c.String("description")),
		Ref:         vela.String(ref),
		Target:      vela.String(target),
		Task:        vela.String(c.String("task")),
	}

	deployment, _, err := client.Deployment.Add(org, repo, request)
	if err != nil {
		return err
	}

	fmt.Printf("deployment \"%s\" was added \n", deployment.GetURL())

	return nil
}
