// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/build"

	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// BuildView defines the command for inspecting a build.
var BuildView = &cli.Command{
	Name:        "build",
	Description: "Use this command to view a build.",
	Usage:       "View details of the provided build",
	Action:      buildView,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG"},
			Name:    "org",
			Usage:   "Provide the organization for the build",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO"},
			Name:    "repo",
			Usage:   "Provide the repository for the build",
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD"},
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Provide the build number",
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
 1. View build details for a repository.
    $ {{.HelpName}} --org github --repo octocat --build-number 1
 2. View build details for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --build-number 1 --output json
 3. View build details for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --build-number 1
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// inspect a build.
func buildView(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the build configuration
	b := &build.Config{
		Action: viewAction,
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Number: c.Int("build"),
		Output: c.String("output"),
	}

	// validate build configuration
	err = b.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the build configuration
	return b.View(client)
}
