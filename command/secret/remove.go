// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// nolint: dupl // ignore similar code among actions
package secret

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/secret"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/go-vela/types/constants"

	"github.com/urfave/cli/v2"
)

// CommandRemove defines the command for inspecting a secret.
var CommandRemove = &cli.Command{
	Name:        "secret",
	Description: "Use this command to remove a secret.",
	Usage:       "Remove details of the provided secret",
	Action:      remove,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "SECRET_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the secret",
			Value:   internal.GetGitConfigOrg("./"),
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "SECRET_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the secret",
			Value:   internal.GetGitConfigRepo("./"),
		},

		// Secret Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ENGINE", "SECRET_ENGINE"},
			Name:    internal.FlagSecretEngine,
			Aliases: []string{"e"},
			Usage:   "provide the engine that stores the secret",
			Value:   constants.DriverNative,
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TYPE", "SECRET_TYPE"},
			Name:    internal.FlagSecretType,
			Aliases: []string{"ty"},
			Usage:   "provide the type of secret being stored",
			Value:   constants.SecretRepo,
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TEAM", "SECRET_TEAM"},
			Name:    "team",
			Aliases: []string{"t"},
			Usage:   "provide the team for the secret",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_NAME", "SECRET_NAME"},
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "provide the name of the secret",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "SECRET_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	// nolint: lll // ignore long line length due to flags
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Remove repository secret details.
    $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo
  2. Remove organization secret details.
    $ {{.HelpName}} --secret.engine native --secret.type org --org MyOrg --name foo
  3. Remove shared secret details.
    $ {{.HelpName}} --secret.engine native --secret.type shared --org MyOrg --team octokitties --name foo
  4. Remove repository secret details with json output.
    $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --output json
  5. Remove secret details when config or environment variables are set.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --name foo

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/secret/remove/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to remove a secret.
func remove(c *cli.Context) error {
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
		Action: internal.ActionRemove,
		Engine: c.String(internal.FlagSecretEngine),
		Type:   c.String(internal.FlagSecretType),
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Team:   c.String("team"),
		Name:   c.String("name"),
		Output: c.String(internal.FlagOutput),
	}

	// validate secret configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the remove call for the secret configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.Remove
	return s.Remove(client)
}
