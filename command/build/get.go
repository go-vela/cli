// SPDX-License-Identifier: Apache-2.0

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

// CommandGet defines the command for capturing a list of builds.
var CommandGet = &cli.Command{
	Name:        internal.FlagBuild,
	Aliases:     []string{"builds"},
	Description: "Use this command to get a list of builds.",
	Usage:       "Display a list of builds",
	Action:      get,
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
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_EVENT", "BUILD_EVENT"),
			Name:    "event",
			Aliases: []string{"e"},
			Usage:   "provide the event filter for the build",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_STATUS", "BUILD_STATUS"),
			Name:    "status",
			Aliases: []string{"s"},
			Usage:   "provide the status filter for the build",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_BRANCH", "BUILD_BRANCH"),
			Name:    "branch",
			Aliases: []string{"b"},
			Usage:   "provide the branch filter for the build",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "BUILD_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},

		// Time Flags

		&cli.Int64Flag{
			Sources: cli.EnvVars("VELA_BEFORE", "BUILD_BEFORE"),
			Name:    internal.FlagBefore,
			Aliases: []string{"bf"},
			Usage:   "before time constraint",
		},
		&cli.Int64Flag{
			Sources: cli.EnvVars("VELA_AFTER", "BUILD_AFTER"),
			Name:    internal.FlagAfter,
			Aliases: []string{"af"},
			Usage:   "after time constraint",
		},

		// Pagination Flags

		&cli.IntFlag{
			Sources: cli.EnvVars("VELA_PAGE", "BUILD_PAGE"),
			Name:    internal.FlagPage,
			Aliases: []string{"p"},
			Usage:   "print a specific page of builds",
			Value:   1,
		},
		&cli.IntFlag{
			Sources: cli.EnvVars("VELA_PER_PAGE", "BUILD_PER_PAGE"),
			Name:    internal.FlagPerPage,
			Aliases: []string{"pp"},
			Usage:   "number of builds to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get builds for a repository.
    $ {{.HelpName}} --org MyOrg --repo MyRepo
  2. Get builds for a repository with the pull_request event.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --event pull_request
  3. Get builds for a repository with the status of success.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --status success
  4. Get builds for a repository with the branch of main.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --branch main
  5. Get builds for a repository with wide view output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --output wide
  6. Get builds for a repository with yaml output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --output yaml
  7. Get builds for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --output json
  8. Get builds for a repository when config or environment variables are set.
    $ {{.HelpName}}
  9. Get builds for a repository that were created before 1/2/22 & after 1/1/22.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --before 1641081600 --after 1640995200

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/build/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to capture a list
// of builds.
func get(_ context.Context, c *cli.Command) error {
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
		Action:  internal.ActionGet,
		Org:     c.String(internal.FlagOrg),
		Repo:    c.String(internal.FlagRepo),
		Event:   c.String("event"),
		Status:  c.String("status"),
		Branch:  c.String("branch"),
		Before:  c.Int64(internal.FlagBefore),
		After:   c.Int64(internal.FlagAfter),
		Page:    c.Int(internal.FlagPage),
		PerPage: c.Int(internal.FlagPerPage),
		Output:  c.String(internal.FlagOutput),
		Color:   output.ColorOptionsFromCLIContext(c),
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
