// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/step"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// StepView defines the command for inspecting a step.
var StepView = &cli.Command{
	Name:        "step",
	Description: "Use this command to view a step.",
	Usage:       "View details of the provided step",
	Action:      stepView,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "STEP_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "provide the organization for the step",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "STEP_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "provide the repository for the step",
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD", "STEP_BUILD"},
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "provide the build for the step",
		},

		// Step Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_STEP", "STEP_NUMBER"},
			Name:    "step",
			Aliases: []string{"s"},
			Usage:   "provide the number for the step",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "STEP_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "print the output in default, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View step details for a repository.
    $ {{.HelpName}} --org github --repo octocat --build 1 --step 1
  2. View step details for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --build 1 --step 1 --output json
  3. View step details for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --build 1 --step 1

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/step/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// inspect a step.
func stepView(c *cli.Context) error {
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
		Action: viewAction,
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Build:  c.Int("build"),
		Number: c.Int("step"),
		Output: c.String("output"),
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
