// SPDX-License-Identifier: Apache-2.0

package step

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/step"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandView defines the command for inspecting a step.
var CommandView = &cli.Command{
	Name:        internal.FlagStep,
	Description: "Use this command to view a step.",
	Usage:       "View details of the provided step",
	Action:      view,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ORG", "STEP_ORG"),
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the step",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REPO", "STEP_REPO"),
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the step",
		},

		// Build Flags

		&cli.Int64Flag{
			Sources: cli.EnvVars("VELA_BUILD", "STEP_BUILD"),
			Name:    internal.FlagBuild,
			Aliases: []string{"b"},
			Usage:   "provide the build for the step",
		},

		// Step Flags

		&cli.Int32Flag{
			Sources: cli.EnvVars("VELA_STEP", "STEP_NUMBER"),
			Name:    internal.FlagStep,
			Aliases: []string{"s", "number", "sn"},
			Usage:   "provide the number for the step",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "STEP_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View step details for a repository.
    $ {{.FullName}} --org MyOrg --repo MyRepo --build 1 --step 1
  2. View step details for a repository with json output.
    $ {{.FullName}} --org MyOrg --repo MyRepo --build 1 --step 1 --output json
  3. View step details for a repository config or environment variables are set.
    $ {{.FullName}} --build 1 --step 1

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/step/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to inspect a step.
func view(ctx context.Context, c *cli.Command) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// grab first command line argument, if it exists, and set it as resource
	err = internal.ProcessArgs(c, internal.FlagStep, "int")
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

	// create the step configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/step?tab=doc#Config
	s := &step.Config{
		Action: internal.ActionView,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Build:  c.Int64(internal.FlagBuild),
		Number: c.Int32(internal.FlagStep),
		Output: c.String(internal.FlagOutput),
		Color:  output.ColorOptionsFromCLIContext(c),
	}

	// validate step configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/step?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the step configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/step?tab=doc#Config.View
	return s.View(ctx, client)
}
