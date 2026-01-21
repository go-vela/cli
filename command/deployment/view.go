// SPDX-License-Identifier: Apache-2.0

package deployment

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/deployment"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandView defines the command for inspecting a deployment.
var CommandView = &cli.Command{
	Name:        "deployment",
	Description: "Use this command to view a deployment.",
	Usage:       "View details of the provided deployment",
	Action:      view,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ORG", "DEPLOYMENT_ORG"),
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the deployment",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REPO", "DEPLOYMENT_REPO"),
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the deployment",
		},

		// Deployment Flags

		&cli.Int64Flag{
			Sources: cli.EnvVars("VELA_DEPLOYMENT", "DEPLOYMENT_NUMBER"),
			Name:    "deployment",
			Aliases: []string{"d", "number", "dn"},
			Usage:   "provide the number for the deployment",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "DEPLOYMENT_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View deployment details for a repository.
    $ {{.FullName}} --org MyOrg --repo MyRepo --deployment 1
  2. View deployment details for a repository with json output.
    $ {{.FullName}} --org MyOrg --repo MyRepo --deployment 1 --output json
  3. View deployment details for a repository config or environment variables are set.
    $ {{.FullName}} --deployment 1

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/deployment/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to inspect a
// deployment.
func view(ctx context.Context, c *cli.Command) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// grab first command line argument, if it exists, and set it as resource
	err = internal.ProcessArgs(c, "deployment", "int")
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
		Action: internal.ActionView,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Number: c.Int64("deployment"),
		Output: c.String(internal.FlagOutput),
		Color:  output.ColorOptionsFromCLIContext(c),
	}

	// validate deployment configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/deployment?tab=doc#Config.Validate
	err = d.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the deployment configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/deployment?tab=doc#Config.View
	return d.View(ctx, client)
}
