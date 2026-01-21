// SPDX-License-Identifier: Apache-2.0

package user

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/user"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
)

// CommandView defines the command for inspecting a user.
var CommandView = &cli.Command{
	Name:        "user",
	Description: "Use this command to view a user.",
	Usage:       "View details of the provided user",
	Action:      view,
	Flags: []cli.Flag{

		// User Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_USER_NAME", "USER_NAME"),
			Name:    internal.FlagName,
			Aliases: []string{"n"},
			Usage:   "provide the name of the user to view",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "REPO_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View details of the current user.
    $ {{.FullName}}
  2. View details of another user (admin).
    $ {{.FullName}} --name Octocat
  3. View details of current user with json output.
    $ {{.FullName}} --output json

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/repo/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to inspect a user.
func view(ctx context.Context, c *cli.Command) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	client, err := client.Parse(c)
	if err != nil {
		return err
	}

	u := &user.Config{
		Name:   c.String(internal.FlagName),
		Output: c.String(internal.FlagOutput),
		Color:  output.ColorOptionsFromCLIContext(c),
	}

	return u.View(ctx, client)
}
