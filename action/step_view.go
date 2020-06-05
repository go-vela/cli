// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/step"

	"github.com/go-vela/sdk-go/vela"

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
			EnvVars: []string{"VELA_ORG"},
			Name:    "org",
			Usage:   "Provide the organization for the step",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO"},
			Name:    "repo",
			Usage:   "Provide the repository for the step",
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD"},
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Provide the build number for the step",
		},

		// Step Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_STEP"},
			Name:    "step",
			Aliases: []string{"s"},
			Usage:   "Provide the step number",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT"},
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Print the output in wide, yaml or json format",
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
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// inspect a step.
func stepView(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the step configuration
	s := &step.Config{
		Action: viewAction,
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Build:  c.Int("build"),
		Number: c.Int("step"),
		Output: c.String("output"),
	}

	// validate step configuration
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the step configuration
	return s.View(client)
}
