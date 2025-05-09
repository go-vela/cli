// SPDX-License-Identifier: Apache-2.0

package log

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/log"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandGet defines the command for capturing a list of build logs.
var CommandGet = &cli.Command{
	Name:        "log",
	Aliases:     []string{"logs"},
	Description: "Use this command to get a list of build logs.",
	Usage:       "Display a list of build logs",
	Action:      get,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ORG", "LOG_ORG"),
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the log",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REPO", "LOG_REPO"),
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the log",
		},

		// Build Flags

		&cli.Int64Flag{
			Sources: cli.EnvVars("VELA_BUILD", "LOG_BUILD"),
			Name:    internal.FlagBuild,
			Aliases: []string{"b"},
			Usage:   "provide the build for the log",
		},

		// Pagination Flags

		&cli.IntFlag{
			Sources: cli.EnvVars("VELA_PAGE", "BUILD_PAGE"),
			Name:    internal.FlagPage,
			Aliases: []string{"p"},
			Usage:   "print a specific page of logs",
			Value:   1,
		},
		&cli.IntFlag{
			Sources: cli.EnvVars("VELA_PER_PAGE", "BUILD_PER_PAGE"),
			Name:    internal.FlagPerPage,
			Aliases: []string{"pp"},
			Usage:   "number of logs to print per page",
			Value:   100,
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "LOG_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get logs for a build.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1
  2. Get logs for a build with yaml output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --output yaml
  3. Get logs for a build with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --output json
  4. Get logs for a build when config or environment variables are set.
    $ {{.HelpName}} --build 1

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/log/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to capture a list
// of build logs.
func get(ctx context.Context, c *cli.Command) error {
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

	// create the log configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/log?tab=doc#Config
	l := &log.Config{
		Action:  internal.ActionGet,
		Org:     c.String(internal.FlagOrg),
		Repo:    c.String(internal.FlagRepo),
		Build:   c.Int64(internal.FlagBuild),
		Page:    c.Int(internal.FlagPage),
		PerPage: c.Int(internal.FlagPerPage),
		Output:  c.String(internal.FlagOutput),
		Color:   output.ColorOptionsFromCLIContext(c),
	}

	// validate log configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/log?tab=doc#Config.Validate
	err = l.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the log configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/log?tab=doc#Config.Get
	return l.Get(client)
}
