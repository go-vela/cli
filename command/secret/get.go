// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

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

// CommandGet defines the command for inspecting a secret.
var CommandGet = &cli.Command{
	Name:        "secret",
	Aliases:     []string{"secrets"},
	Description: "Use this command to get a list of secrets.",
	Usage:       "Display a list of secrets",
	Action:      get,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "SECRET_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the secret",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "SECRET_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the secret",
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

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "SECRET_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},

		// Pagination Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_PAGE", "SECRET_PAGE"},
			Name:    internal.FlagPage,
			Aliases: []string{"p"},
			Usage:   "print a specific page of secrets",
			Value:   1,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_PER_PAGE", "SECRET_PER_PAGE"},
			Name:    internal.FlagPerPage,
			Aliases: []string{"pp"},
			Usage:   "number of secrets to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get repository secret details.
    $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo
  2. Get organization secret details.
    $ {{.HelpName}} --secret.engine native --secret.type org --org MyOrg
  3. Get shared secret details.
    $ {{.HelpName}} --secret.engine native --secret.type shared --org MyOrg --team octokitties
  4. Get repository secret details with json output.
    $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --output json
  5. Get secret details when config or environment variables are set.
    $ {{.HelpName}} --org MyOrg --repo MyRepo

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/secret/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to capture a list
// of secrets.
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

	// create the secret configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config
	s := &secret.Config{
		Action:  internal.ActionGet,
		Engine:  c.String(internal.FlagSecretEngine),
		Type:    c.String(internal.FlagSecretType),
		Org:     c.String(internal.FlagOrg),
		Repo:    c.String(internal.FlagRepo),
		Team:    c.String("team"),
		Page:    c.Int(internal.FlagPage),
		PerPage: c.Int(internal.FlagPerPage),
		Output:  c.String(internal.FlagOutput),
	}

	// validate secret configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the secret configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.Get
	return s.Get(client)
}
