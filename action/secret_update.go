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

// SecretUpdate defines the command for inspecting a secret.
var SecretUpdate = &cli.Command{
	Name:        "secret",
	Description: "Use this command to view a secret.",
	Usage:       "Update details of the provided secret",
	Action:      secretUpdate,
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
		&cli.StringFlag{
			EnvVars: []string{"VELA_NAME", "SECRET_NAME"},
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "Provide the name of the secret",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_VALUE", "SECRET_VALUE"},
			Name:    "value",
			Aliases: []string{"v"},
			Usage:   "Provide the value for the secret",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_IMAGES", "SECRET_IMAGES"},
			Name:    "image",
			Aliases: []string{"i"},
			Usage:   "Provide the image(s) that can access this secret",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_EVENTS", "SECRET_EVENTS"},
			Name:    "event",
			Aliases: []string{"ev"},
			Usage:   "Provide the event(s) that can access this secret",
			Value: cli.NewStringSlice(
				constants.EventDeploy,
				constants.EventPush,
				constants.EventTag,
			),
		},
		&cli.BoolFlag{
			EnvVars: []string{"VELA_COMMAND", "SECRET_COMMAND"},
			Name:    "commands",
			Aliases: []string{"c"},
			Usage:   "Enable a secret to be used for a step with commands",
			Value:   true,
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_FILE", "SECRET_FILE"},
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "Provide a file to update the secret(s)",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "SECRET_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "Print the output in default, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Update a repository secret.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value bar
  2. Update an organization secret.
    $ {{.HelpName}} --engine native --type org --org github --name foo --value bar
  3. Update a shared secret.
    $ {{.HelpName}} --engine native --type shared --org github --team octokitties --name foo --value bar
  4. Update a repository secret with all event types enabled.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --event comment --event deployment --event pull_request --event push --event tag
  5. Update a repository secret with an image whitelist.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --image alpine --image golang:* --image postgres:latest
  6. Update a secret with value from a file.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value @secret.txt
  7. Update a repository secret with json output.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value bar --output json
  8. Update a secret or secrets from a file.
    $ {{.HelpName}} --file secret.yml
  9. Update a secret when engine and type config or environment variables are set.
    $ {{.HelpName}} --org github --repo octocat --name foo --value bar
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// modify a secret.
func secretUpdate(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the secret configuration
	s := &secret.Config{
		Action:       addAction,
		Engine:       c.String("engine"),
		Type:         c.String("type"),
		Org:          c.String("org"),
		Repo:         c.String("repo"),
		Team:         c.String("team"),
		Name:         c.String("name"),
		Value:        c.String("value"),
		AllowCommand: c.Bool("commands"),
		Images:       c.StringSlice("image"),
		Events:       c.StringSlice("event"),
		File:         c.String("file"),
		Output:       c.String("output"),
	}

	// validate secret configuration
	err = s.Validate()
	if err != nil {
		return err
	}

	// check if secret file is provided
	if len(s.File) > 0 {
		// execute the update from file call for the secret configuration
		return s.UpdateFromFile(client)
	}

	// execute the update call for the secret configuration
	return s.Update(client)
}
