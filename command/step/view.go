// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package step

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/step"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
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
			EnvVars: []string{"VELA_ORG", "STEP_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the step",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "STEP_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the step",
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD", "STEP_BUILD"},
			Name:    internal.FlagBuild,
			Aliases: []string{"b"},
			Usage:   "provide the build for the step",
		},

		// Step Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_STEP", "STEP_NUMBER"},
			Name:    internal.FlagStep,
			Aliases: []string{"s", "number", "sn"},
			Usage:   "provide the number for the step",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "STEP_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View step details for a repository.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --step 1
  2. View step details for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --step 1 --output json
  3. View step details for a repository config or environment variables are set.
    $ {{.HelpName}} --build 1 --step 1

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/step/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to inspect a step.
func view(c *cli.Context) error {
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
		Build:  c.Int(internal.FlagBuild),
		Number: c.Int(internal.FlagStep),
		Output: c.String(internal.FlagOutput),
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
	return s.View(client)
}
