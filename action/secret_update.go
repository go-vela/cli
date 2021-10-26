// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/secret"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/go-vela/types/constants"

	"github.com/urfave/cli/v2"
)

// SecretUpdate defines the command for inspecting a secret.
var SecretUpdate = &cli.Command{
	Name:        "secret",
	Description: "Use this command to update a secret.",
	Usage:       "Update details of the provided secret",
	Action:      secretUpdate,
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
		&cli.StringFlag{
			EnvVars: []string{"VELA_NAME", "SECRET_NAME"},
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "provide the name of the secret",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_VALUE", "SECRET_VALUE"},
			Name:    "value",
			Aliases: []string{"v"},
			Usage:   "provide the value for the secret",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_IMAGES", "SECRET_IMAGES"},
			Name:    "image",
			Aliases: []string{"i"},
			Usage:   "provide the image(s) that can access this secret",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_EVENTS", "SECRET_EVENTS"},
			Name:    "event",
			Aliases: []string{"ev"},
			Usage:   "provide the event(s) that can access this secret",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_COMMAND", "SECRET_COMMAND"},
			Name:    "commands",
			Aliases: []string{"c"},
			Usage:   "enable a secret to be used for a step with commands",
			Value:   "true",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_FILE", "SECRET_FILE"},
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "provide a file to update the secret(s)",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "SECRET_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "Print the output in default, yaml or json format",
		},
	},
	// nolint: lll // ignore long line length due to flags
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Update a repository secret.
    $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --value bar
  2. Update an organization secret.
    $ {{.HelpName}} --secret.engine native --secret.type org --org MyOrg --name foo --value bar
  3. Update a shared secret.
    $ {{.HelpName}} --secret.engine native --secret.type shared --org MyOrg --team octokitties --name foo --value bar
  4. Update a repository secret with all event types enabled.
    $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --event comment --event deployment --event pull_request --event push --event tag
  5. Update a repository secret with an image whitelist.
    $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --image alpine --image golang:* --image postgres:latest
  6. Update a secret with value from a file.
    $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --value @secret.txt
  7. Update a repository secret with json output.
    $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --value bar --output json
  8. Update a secret or secrets from a file.
    $ {{.HelpName}} --file secret.yml
  9. Update a secret when config or environment variables are set.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --name foo --value bar

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/secret/update/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// modify a secret.
//
// nolint: dupl // ignore similar code among actions
func secretUpdate(c *cli.Context) error {
	// load variables from the config file
	err := Load(c)
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
		Action:       updateAction,
		Engine:       c.String(internal.FlagSecretEngine),
		Type:         c.String(internal.FlagSecretType),
		Org:          c.String(internal.FlagOrg),
		Repo:         c.String(internal.FlagRepo),
		Team:         c.String("team"),
		Name:         c.String("name"),
		Value:        c.String("value"),
		AllowCommand: c.Bool("commands"),
		Images:       c.StringSlice("image"),
		Events:       c.StringSlice("event"),
		File:         c.String("file"),
		Output:       c.String(internal.FlagOutput),
	}

	// validate secret configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// check if secret file is provided
	if len(s.File) > 0 {
		// execute the update from file call for the secret configuration
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.UpdateFromFile
		return s.UpdateFromFile(client)
	}

	// execute the update call for the secret configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.Update
	return s.Update(client)
}
