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

// BuildView defines the command for inspecting a build.
var BuildView = &cli.Command{
	Name:        "build",
	Description: "Use this command to view a build.",
	Usage:       "View details of the provided build",
	Action:      buildView,
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

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD", "BUILD_NUMBER"},
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "provide the number for the build",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "BUILD_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View build details for a repository.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --build 1
  2. View build details for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --build 1 --output json
  3. View build details for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --build 1

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/build/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// inspect a build.
func buildView(c *cli.Context) error {
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
		Action: viewAction,
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Number: c.Int("build"),
		Output: c.String("output"),
	}

	// validate build configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/build?tab=doc#Config.Validate
	err = b.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the build configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/build?tab=doc#Config.View
	return b.View(client)
}
