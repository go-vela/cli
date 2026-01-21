// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/repo"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandGet defines the command for capturing a list of repositories.
var CommandGet = &cli.Command{
	Name:        "repo",
	Aliases:     []string{"repos"},
	Description: "Use this command to get a list of repositories.",
	Usage:       "Display a list of repositories",
	Action:      get,
	Flags: []cli.Flag{

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "REPO_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},

		// Pagination Flags

		&cli.IntFlag{
			Sources: cli.EnvVars("VELA_PAGE", "REPO_PAGE"),
			Name:    internal.FlagPage,
			Aliases: []string{"p"},
			Usage:   "print a specific page of repositories",
			Value:   1,
		},
		&cli.IntFlag{
			Sources: cli.EnvVars("VELA_PER_PAGE", "REPO_PER_PAGE"),
			Name:    internal.FlagPerPage,
			Aliases: []string{"pp"},
			Usage:   "number of repositories to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get a list of repositories.
    $ {{.FullName}}
  2. Get a list of repositories with wide view output.
    $ {{.FullName}} --output wide
  3. Get a list of repositories with yaml output.
    $ {{.FullName}} --output yaml
  4. Get a list of repositories with json output.
    $ {{.FullName}} --output json
  5. Get a list of repositories when config or environment variables are set.
    $ {{.FullName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/repo/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to capture a list
// of repos.
//
//nolint:dupl // duplicate of `command/repo/remove.go:70-108`
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

	// create the repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config
	r := &repo.Config{
		Action:  internal.ActionGet,
		Page:    c.Int(internal.FlagPage),
		PerPage: c.Int(internal.FlagPerPage),
		Output:  c.String(internal.FlagOutput),
		Color:   output.ColorOptionsFromCLIContext(c),
	}

	// validate repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Validate
	err = r.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Get
	return r.Get(ctx, client)
}
