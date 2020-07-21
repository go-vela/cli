// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/service"
	"github.com/go-vela/cli/internal/client"

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
			EnvVars: []string{"VELA_ORG", "SERVICE_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "provide the organization for the build",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "SERVICE_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "provide the repository for the build",
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD", "SERVICE_BUILD"},
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "provide the build for the service",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "SERVICE_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},

		// Pagination Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_PAGE", "SERVICE_PAGE"},
			Name:    "page",
			Aliases: []string{"p"},
			Usage:   "print a specific page of services",
			Value:   1,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_PER_PAGE", "SERVICE_PER_PAGE"},
			Name:    "per.page",
			Aliases: []string{"pp"},
			Usage:   "number of services to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get services for a repository.
    $ {{.HelpName}} --org MyOrg --repo octocat --build 1
  2. Get services for a repository with wide view output.
    $ {{.HelpName}} --org MyOrg --repo octocat --build 1 --output wide
  3. Get services for a repository with yaml output.
    $ {{.HelpName}} --org MyOrg --repo octocat --build 1 --output yaml
  4. Get services for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo octocat --build 1 --output json
  5. Get services for a build when config or environment variables are set.
    $ {{.HelpName}} --build 1

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/service/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// capture a list of services.
func serviceGet(c *cli.Context) error {
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
		Action:  getAction,
		Org:     c.String("org"),
		Repo:    c.String("repo"),
		Build:   c.Int("build"),
		Page:    c.Int("page"),
		PerPage: c.Int("per.page"),
		Output:  c.String("output"),
	}

	// validate service configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/service?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the service configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/service?tab=doc#Config.Get
	return s.Get(client)
}
