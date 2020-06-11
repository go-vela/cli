// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/secret"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/constants"

	"github.com/urfave/cli/v2"
)

// SecretRemove defines the command for inspecting a secret.
var SecretRemove = &cli.Command{
	Name:        "secret",
	Description: "Use this command to view a secret.",
	Usage:       "Remove details of the provided secret",
	Action:      secretRemove,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "SECRET_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "provide the organization for the secret",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "SECRET_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "provide the repository for the secret",
		},

		// Secret Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ENGINE", "SECRET_ENGINE"},
			Name:    "engine",
			Aliases: []string{"e"},
			Usage:   "provide the engine that stores the secret",
			Value:   constants.DriverNative,
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TYPE", "SECRET_TYPE"},
			Name:    "type",
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
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "print the output in default, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Remove repository secret details.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo
  2. Remove organization secret details.
    $ {{.HelpName}} --engine native --type org --org github --name foo
  3. Remove shared secret details.
    $ {{.HelpName}} --engine native --type shared --org github --team octokitties --name foo
  4. Remove repository secret details with json output.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --output json
  5. Remove secret details when engine and type config or environment variables are set.
    $ {{.HelpName}} --org github --repo octocat --name foo

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/secret/remove/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// remove a secret.
func secretRemove(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the secret configuration
	s := &secret.Config{
		Action: removeAction,
		Engine: c.String("engine"),
		Type:   c.String("type"),
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Team:   c.String("team"),
		Name:   c.String("name"),
		Output: c.String("output"),
	}

	// validate secret configuration
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the remove call for the secret configuration
	return s.Remove(client)
}