// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code among actions
package secret

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/secret"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/server/constants"
)

// CommandView defines the command for inspecting a secret.
var CommandView = &cli.Command{
	Name:        "secret",
	Description: "Use this command to view a secret.",
	Usage:       "View details of the provided secret",
	Action:      view,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ORG", "SECRET_ORG"),
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the secret",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_REPO", "SECRET_REPO"),
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the secret",
		},

		// Secret Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_ENGINE", "SECRET_ENGINE"),
			Name:    internal.FlagSecretEngine,
			Aliases: []string{"e"},
			Usage:   "provide the engine that stores the secret",
			Value:   constants.DriverNative,
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_TYPE", "SECRET_TYPE"),
			Name:    internal.FlagSecretType,
			Aliases: []string{"ty"},
			Usage:   "provide the type of secret being stored",
			Value:   constants.SecretRepo,
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_TEAM", "SECRET_TEAM"),
			Name:    "team",
			Aliases: []string{"t"},
			Usage:   "provide the team for the secret",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_NAME", "SECRET_NAME"),
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "provide the name of the secret",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "SECRET_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View repository secret details.
    $ {{.FullName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo
  2. View organization secret details.
    $ {{.FullName}} --secret.engine native --secret.type org --org MyOrg --name foo
  3. View shared secret details.
    $ {{.FullName}} --secret.engine native --secret.type shared --org MyOrg --team octokitties --name foo
  4. View repository secret details with json output.
    $ {{.FullName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --output json
  5. View secret details when config or environment variables are set.
    $ {{.FullName}} --org MyOrg --repo MyRepo --name foo

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/secret/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// inspect a secret.
func view(_ context.Context, c *cli.Command) error {
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

	// create the secret configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config
	s := &secret.Config{
		Action: internal.ActionView,
		Engine: c.String(internal.FlagSecretEngine),
		Type:   c.String(internal.FlagSecretType),
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Team:   c.String("team"),
		Name:   c.String("name"),
		Output: c.String(internal.FlagOutput),
		Color:  output.ColorOptionsFromCLIContext(c),
	}

	// validate secret configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the secret configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.View
	return s.View(client)
}
