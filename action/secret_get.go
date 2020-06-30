// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/secret"
	"github.com/go-vela/cli/internal/client"

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

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "SECRET_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "print the output in default, wide, yaml or json format",
		},

		// Pagination Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_PAGE", "SECRET_PAGE"},
			Name:    "page",
			Aliases: []string{"p"},
			Usage:   "print a specific page of secrets",
			Value:   1,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_PER_PAGE", "SECRET_PER_PAGE"},
			Name:    "per.page",
			Aliases: []string{"pp"},
			Usage:   "number of secrets to print per page",
			Value:   10,
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

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/secret/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// capture a list of secrets.
func secretGet(c *cli.Context) error {
	// parse the Vela client from the context
	client, err := client.Parse(c)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the secret configuration
	s := &secret.Config{
		Action:  getAction,
		Engine:  c.String("engine"),
		Type:    c.String("type"),
		Org:     c.String("org"),
		Repo:    c.String("repo"),
		Team:    c.String("team"),
		Page:    c.Int("page"),
		PerPage: c.Int("per.page"),
		Output:  c.String("output"),
	}

	// validate secret configuration
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the secret configuration
	return s.Get(client)
}
