// SPDX-License-Identifier: Apache-2.0

package step

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/step"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// CommandGet defines the command for capturing a list of steps.
var CommandGet = &cli.Command{
	Name:        internal.FlagStep,
	Aliases:     []string{"steps"},
	Description: "Use this command to get a list of steps.",
	Usage:       "Display a list of steps",
	Action:      get,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "STEP_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the step",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "STEP_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the step",
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD", "STEP_BUILD"},
			Name:    internal.FlagBuild,
			Aliases: []string{"b"},
			Usage:   "provide the build for the step",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "STEP_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},

		// Pagination Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_PAGE"},
			Name:    internal.FlagPage,
			Aliases: []string{"p"},
			Usage:   "print a specific page of steps",
			Value:   1,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_PER_PAGE"},
			Name:    internal.FlagPerPage,
			Aliases: []string{"pp"},
			Usage:   "number of steps to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get steps for a repository.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1
  2. Get steps for a repository with wide view output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --output wide
  3. Get steps for a repository with yaml output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --output yaml
  4. Get steps for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --output json
  5. Get steps for a build when config or environment variables are set.
    $ {{.HelpName}} --build 1

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/step/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to capture a list
// of steps.
func get(c *cli.Context) error {
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

	// create the step configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/step?tab=doc#Config
	s := &step.Config{
		Action:  internal.ActionGet,
		Org:     c.String(internal.FlagOrg),
		Repo:    c.String(internal.FlagRepo),
		Build:   c.Int(internal.FlagBuild),
		Page:    c.Int(internal.FlagPage),
		PerPage: c.Int(internal.FlagPerPage),
		Output:  c.String(internal.FlagOutput),
	}

	// validate step configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/step?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the step configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/step?tab=doc#Config.Get
	return s.Get(client)
}
