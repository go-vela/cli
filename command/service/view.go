// SPDX-License-Identifier: Apache-2.0

package service

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/service"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandView defines the command for inspecting a service.
var CommandView = &cli.Command{
	Name:        internal.FlagService,
	Description: "Use this command to view a service.",
	Usage:       "View details of the provided service",
	Action:      view,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ORG", "SERVICE_ORG"),
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the service",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REPO", "SERVICE_REPO"),
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the service",
		},

		// Build Flags

		&cli.Int64Flag{
			Sources: cli.EnvVars("VELA_BUILD", "SERVICE_BUILD"),
			Name:    internal.FlagBuild,
			Aliases: []string{"b"},
			Usage:   "provide the build for the service",
		},

		// Service Flags

		&cli.Int32Flag{
			Sources: cli.EnvVars("VELA_SERVICE", "SERVICE_NUMBER"),
			Name:    internal.FlagService,
			Aliases: []string{"s", "number", "sn"},
			Usage:   "provide the number for the service",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "SERVICE_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View service details for a repository.
    $ {{.FullName}} --org MyOrg --repo MyRepo --build 1 --service 1
  2. View service details for a repository with json output.
    $ {{.FullName}} --org MyOrg --repo MyRepo --build 1 --service 1 --output json
  3. View service details for a repository when config or environment variables are set.
    $ {{.FullName}} --build 1 --service 1

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/service/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to inspect a service.
func view(_ context.Context, c *cli.Command) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// grab first command line argument, if it exists, and set it as resource
	err = internal.ProcessArgs(c, internal.FlagService, "int")
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

	// create the service configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/service?tab=doc#Config
	s := &service.Config{
		Action: internal.ActionView,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Build:  c.Int64(internal.FlagBuild),
		Number: c.Int32(internal.FlagService),
		Output: c.String(internal.FlagOutput),
		Color:  output.ColorOptionsFromCLIContext(c),
	}

	// validate service configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/service?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the service configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/service?tab=doc#Config.View
	return s.View(client)
}
