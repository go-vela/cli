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

// StepGet defines the command for capturing a list of steps.
var StepGet = &cli.Command{
	Name:        "step",
	Aliases:     []string{"steps"},
	Description: "Use this command to get a list of steps.",
	Usage:       "Display a list of steps",
	Action:      stepGet,
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

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT"},
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Print the output in wide, yaml or json format",
		},

		// Pagination Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_PAGE"},
			Name:    "page",
			Aliases: []string{"p"},
			Usage:   "Print a specific page of steps",
			Value:   1,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_PER_PAGE"},
			Name:    "per.page",
			Aliases: []string{"pp"},
			Usage:   "Expand the number of items contained within page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Get steps for a repository.
    $ {{.HelpName}} --org github --repo octocat --build 1
 2. Get steps for a repository with wide view output.
    $ {{.HelpName}} --org github --repo octocat --build 1 --output wide
 3. Get steps for a repository with yaml output.
    $ {{.HelpName}} --org github --repo octocat --build 1 --output yaml
 4. Get steps for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --build 1 --output json
 5. Get steps for a build when org and repo config or environment variables are set.
    $ {{.HelpName}} --build 1
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// capture a list of steps.
func stepGet(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the step configuration
	s := &step.Config{
		Action:  getAction,
		Org:     c.String("org"),
		Repo:    c.String("repo"),
		Build:   c.Int("build"),
		Page:    c.Int("page"),
		PerPage: c.Int("per.page"),
		Output:  c.String("output"),
	}

	// validate step configuration
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the step configuration
	return s.Get(client)
}
