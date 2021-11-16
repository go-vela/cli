// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/deployment"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// CommandAdd defines the command for creating a deployment.
var CommandAdd = &cli.Command{
	Name:        "deployment",
	Description: "Use this command to add a deployment.",
	Usage:       "Add a new deployment from the provided configuration",
	Action:      add,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "DEPLOYMENT_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the deployment",
			Value:   internal.GetCwdOrg("./"),
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "DEPLOYMENT_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the deployment",
			Value:   internal.GetCwdRepo("./"),
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
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_PARAMETERS", "DEPLOYMENT_PARAMETERS"},
			Name:    "parameter",
			Aliases: []string{"p"},
			Usage:   "provide the parameter(s) within `key=value` format for the deployment",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "DEPLOYMENT_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Add a deployment for a repository.
    $ {{.HelpName}} --org MyOrg --repo MyRepo
  2. Add a deployment for a repository with a specific target environment.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --target stage
  3. Add a deployment for a repository with a specific branch reference.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --ref dev
  4. Add a deployment for a repository with a specific commit reference.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --ref 48afb5bdc41ad69bf22588491333f7cf71135163
  5. Add a deployment for a repository with a specific tag reference.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --ref refs/tags/1.0.0
  6. Add a deployment for a repository with a specific description.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --description 'my custom message'
  7. Add a deployment for a repository with two parameters.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --parameter 'key=value' --parameter 'foo=bar'
  8. Add a deployment for a repository when config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/deployment/add/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to create a
// deployment.
func add(c *cli.Context) error {
	// load variables from the config file
	err := action.Load(c)
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

	// create the deployment configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/deployment?tab=doc#Config
	d := &deployment.Config{
		Action:      internal.ActionAdd,
		Org:         c.String(internal.FlagOrg),
		Repo:        c.String(internal.FlagRepo),
		Description: c.String("description"),
		Ref:         c.String("ref"),
		Target:      c.String("target"),
		Task:        c.String("task"),
		Output:      c.String(internal.FlagOutput),
		Parameters:  c.StringSlice("parameter"),
	}

	// validate deployment configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/deployment?tab=doc#Config.Validate
	err = d.Validate()
	if err != nil {
		return err
	}

	// execute the add call for the deployment configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/deployment?tab=doc#Config.Add
	return d.Add(client)
}
