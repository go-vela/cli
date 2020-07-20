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

// SecretAdd defines the command for inspecting a secret.
var SecretAdd = &cli.Command{
	Name:        "secret",
	Description: "Use this command to view a secret.",
	Usage:       "Add details of the provided secret",
	Action:      secretAdd,
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
			Usage:   "Provide the image(s) that can access this secret",
		},
		&cli.StringSliceFlag{
			EnvVars: []string{"VELA_EVENTS", "SECRET_EVENTS"},
			Name:    "event",
			Aliases: []string{"ev"},
			Usage:   "provide the event(s) that can access this secret",
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
			Usage:   "enable a secret to be used for a step with commands",
			Value:   true,
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_FILE", "SECRET_FILE"},
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "provide a file to add the secret(s)",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "SECRET_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Add a repository secret.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value bar
  2. Add an organization secret.
    $ {{.HelpName}} --engine native --type org --org github --name foo --value bar
  3. Add a shared secret.
    $ {{.HelpName}} --engine native --type shared --org github --team octokitties --name foo --value bar
  4. Add a repository secret with all event types enabled.
     $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value bar --event comment --event deployment --event pull_request --event push --event tag
  5. Add a repository secret with an image whitelist.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value bar --image alpine --image golang:* --image postgres:latest
  6. Add a secret with value from a file.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value @secret.txt
  7. Add a repository secret with json output.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value bar --output json
  8. Add a secret or secrets from a file.
    $ {{.HelpName}} --file secret.yml
  9. Add a secret when engine and type config or environment variables are set.
    $ {{.HelpName}} --org github --repo octocat --name foo --value bar

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/secret/add/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// create a secret.
func secretAdd(c *cli.Context) error {
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
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.Validate
	err = s.Validate()
	if err != nil {
		return err
	}

	// check if secret file is provided
	if len(s.File) > 0 {
		// execute the add from file call for the secret configuration
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.AddFromFile
		return s.AddFromFile(client)
	}

	// execute the add call for the secret configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.Add
	return s.Add(client)
}
