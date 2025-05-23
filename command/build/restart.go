// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code with cancel
package build

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/build"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandRestart defines the command for restarting a build.
var CommandRestart = &cli.Command{
	Name:        internal.FlagBuild,
	Description: "Use this command to restart a build.",
	Usage:       "Restart the provided build",
	Action:      restart,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ORG", "BUILD_ORG"),
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the build",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REPO", "BUILD_REPO"),
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the build",
		},

		// Build Flags

		&cli.Int64Flag{
			Sources: cli.EnvVars("VELA_BUILD", "BUILD_NUMBER"),
			Name:    internal.FlagBuild,
			Aliases: []string{"b", "number", "bn"},
			Usage:   "provide the number for the build",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "BUILD_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Restart existing build for a repository.
    $ {{.FullName}} --org MyOrg --repo MyRepo --build 1
  2. Restart existing build for a repository when config or environment variables are set.
    $ {{.FullName}} --build 1

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/build/restart/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// restart a build.
func restart(_ context.Context, c *cli.Command) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// grab first command line argument, if it exists, and set it as resource
	err = internal.ProcessArgs(c, internal.FlagBuild, "int")
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
		Action: internal.ActionRestart,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Number: c.Int64(internal.FlagBuild),
		Output: c.String(internal.FlagOutput),
		Color:  output.ColorOptionsFromCLIContext(c),
	}

	// validate build configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/build?tab=doc#Config.Validate
	err = b.Validate()
	if err != nil {
		return err
	}

	// execute the restart call for the build configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/build?tab=doc#Config.Restart
	return b.Restart(client)
}
