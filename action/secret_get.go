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

// SecretGet defines the command for inspecting a secret.
var SecretGet = &cli.Command{
	Name:        "secret",
	Description: "Use this command to get a list of secrets.",
	Usage:       "Display a list of secrets",
	Action:      secretGet,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "SECRET_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "Provide the organization for the secret",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "SECRET_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "Provide the repository for the secret",
		},

		// Secret Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ENGINE", "SECRET_ENGINE"},
			Name:    "engine",
			Aliases: []string{"e"},
			Usage:   "Provide the engine that stores the secret",
			Value:   constants.DriverNative,
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TYPE", "SECRET_TYPE"},
			Name:    "type",
			Aliases: []string{"ty"},
			Usage:   "Provide the type of secret being stored",
			Value:   constants.SecretRepo,
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_TEAM", "SECRET_TEAM"},
			Name:    "team",
			Aliases: []string{"t"},
			Usage:   "Provide the team for the secret",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "SECRET_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "Print the output in default, wide, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get repository secret details.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat
  2. Get organization secret details.
    $ {{.HelpName}} --engine native --type org --org github
  3. Get shared secret details.
    $ {{.HelpName}} --engine native --type shared --org github --team octokitties
  4. Get repository secret details with json output.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --output json
  5. Get secret details when engine and type config or environment variables are set.
    $ {{.HelpName}} --org github --repo octocat
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// capture a list of secrets.
func secretGet(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the secret configuration
	s := &secret.Config{
		Action: getAction,
		Engine: c.String("engine"),
		Type:   c.String("type"),
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Team:   c.String("team"),
		Output: c.String("output"),
	}

	// validate secret configuration
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the secret configuration
	return s.Get(client)
}
