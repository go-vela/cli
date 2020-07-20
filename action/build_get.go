// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/build"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// BuildGet defines the command for capturing a list of builds.
var BuildGet = &cli.Command{
	Name:        "build",
	Aliases:     []string{"builds"},
	Description: "Use this command to get a list of builds.",
	Usage:       "Display a list of builds",
	Action:      buildGet,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "BUILD_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "provide the organization for the build",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "BUILD_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "provide the repository for the build",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "BUILD_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},

		// Pagination Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_PAGE", "BUILD_PAGE"},
			Name:    "page",
			Aliases: []string{"p"},
			Usage:   "print a specific page of builds",
			Value:   1,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_PER_PAGE", "BUILD_PER_PAGE"},
			Name:    "per.page",
			Aliases: []string{"pp"},
			Usage:   "number of builds to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get builds for a repository.
    $ {{.HelpName}} --org MyOrg --repo MyRepo
  2. Get builds for a repository with wide view output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --output wide
  3. Get builds for a repository with yaml output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --output yaml
  4. Get builds for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --output json
  5. Get builds for a repository when config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/build/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// capture a list of builds.
func buildGet(c *cli.Context) error {
	// parse the Vela client from the context
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/client?tab=doc#Parse
	client, err := client.Parse(c)
	if err != nil {
		return err
	}

	// create the build configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/build?tab=doc#Config
	b := &build.Config{
		Action:  getAction,
		Org:     c.String("org"),
		Repo:    c.String("repo"),
		Page:    c.Int("page"),
		PerPage: c.Int("per.page"),
		Output:  c.String("output"),
	}

	// validate build configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/build?tab=doc#Config.Validate
	err = b.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the build configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/build?tab=doc#Config.Get
	return b.Get(client)
}
