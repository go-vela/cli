// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/build"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// CommandCancel defines the command for canceling a build.
var CommandCancel = &cli.Command{
	Name:        internal.FlagBuild,
	Description: "Use this command to cancel a build.",
	Usage:       "Cancel the provided build",
	Action:      cancel,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "BUILD_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the build",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "BUILD_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the build",
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD", "BUILD_NUMBER"},
			Name:    internal.FlagBuild,
			Aliases: []string{"b"},
			Usage:   "provide the number for the build",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "BUILD_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Cancel existing build for a repository.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1
  2. Cancel existing build for a repository when config or environment variables are set.
    $ {{.HelpName}} --build 1

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/build/cancel/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to cancel a build.
func cancel(c *cli.Context) error {
	// load variables from the config file
	err := action.Load(c)
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

	// create the build configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/build?tab=doc#Config
	b := &build.Config{
		Action: internal.ActionCancel,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Number: c.Int(internal.FlagBuild),
		Output: c.String(internal.FlagOutput),
	}

	// validate build configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/build?tab=doc#Config.Validate
	err = b.Validate()
	if err != nil {
		return err
	}

	// execute the cancel call for the build configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/build?tab=doc#Config.Cancel
	return b.Cancel(client)
}
