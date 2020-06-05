// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/service"

	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// ServiceView defines the command for inspecting a service.
var ServiceView = &cli.Command{
	Name:        "service",
	Description: "Use this command to view a service.",
	Usage:       "View details of the provided service",
	Action:      serviceView,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG"},
			Name:    "org",
			Usage:   "Provide the organization for the service",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO"},
			Name:    "repo",
			Usage:   "Provide the repository for the service",
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD"},
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Provide the build number for the service",
		},

		// Service Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_SERVICE"},
			Name:    "service",
			Aliases: []string{"s"},
			Usage:   "Provide the service number",
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
 1. View service details for a repository.
    $ {{.HelpName}} --org github --repo octocat --build 1 --service 1
 2. View service details for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --build 1 --service 1 --output json
 3. View service details for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --build 1 --service 1
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// inspect a service.
func serviceView(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the service configuration
	s := &service.Config{
		Action: viewAction,
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Build:  c.Int("build"),
		Number: c.Int("service"),
		Output: c.String("output"),
	}

	// validate service configuration
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the service configuration
	return s.View(client)
}
