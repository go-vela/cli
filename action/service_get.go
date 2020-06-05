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

// ServiceGet defines the command for capturing a list of services.
var ServiceGet = &cli.Command{
	Name:        "service",
	Aliases:     []string{"services"},
	Description: "Use this command to get a list of services.",
	Usage:       "Display a list of services",
	Action:      serviceGet,
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
			Usage:   "Print a specific page of services",
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
 1. Get services for a repository.
    $ {{.HelpName}} --org github --repo octocat --build 1
 2. Get services for a repository with wide view output.
    $ {{.HelpName}} --org github --repo octocat --build 1 --output wide
 3. Get services for a repository with yaml output.
    $ {{.HelpName}} --org github --repo octocat --build 1 --output yaml
 4. Get services for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --build 1 --output json
 5. Get services for a build when org and repo config or environment variables are set.
    $ {{.HelpName}} --build 1
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// capture a list of services.
func serviceGet(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the service configuration
	s := &service.Config{
		Action:  getAction,
		Org:     c.String("org"),
		Repo:    c.String("repo"),
		Build:   c.Int("build"),
		Page:    c.Int("page"),
		PerPage: c.Int("per.page"),
		Output:  c.String("output"),
	}

	// validate service configuration
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the service configuration
	return s.Get(client)
}
