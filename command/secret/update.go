// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"context"
	"fmt"
	"slices"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/secret"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"
	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/server/constants"
)

// CommandUpdate defines the command for updating a secret.
var CommandUpdate = &cli.Command{
	Name:        "secret",
	Description: "Use this command to update a secret.",
	Usage:       "Update details of the provided secret",
	Action:      update,
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
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_VALUE", "SECRET_VALUE"),
			Name:    "value",
			Aliases: []string{"v"},
			Usage:   "provide the value for the secret",
		},
		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_IMAGES", "SECRET_IMAGES"),
			Name:    "image",
			Aliases: []string{"i"},
			Usage:   "provide the image(s) that can access this secret",
		},
		&cli.StringSliceFlag{
			Sources: cli.EnvVars("VELA_EVENTS", "SECRET_EVENTS"),
			Name:    "event",
			Aliases: []string{"events", "ev"},
			Usage:   "provide the event(s) that can access this secret",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_COMMAND", "SECRET_COMMAND"),
			Name:    internal.FlagSecretCommands,
			Aliases: []string{"c"},
			Usage:   "enable a secret to be used for a step with commands",
			Value:   "true",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_SUBSTITUTION", "SECRET_SUBSTITUTION"),
			Name:    internal.FlagSecretSubstitution,
			Aliases: []string{"s"},
			Usage:   "enable a secret to be substituted",
			Value:   "true",
		},
		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_FILE", "SECRET_FILE"),
			Name:    "file",
			Aliases: []string{"f"},
			Usage:   "provide a file to update the secret(s)",
		},

		// Output Flags

		&cli.StringFlag{
			Sources: cli.EnvVars("VELA_OUTPUT", "SECRET_OUTPUT"),
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "Print the output in default, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
   1. Update a repository secret.
     $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --value bar
   2. Update a repository secret and disallow usage in commands.
     $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --value bar --commands false
   3. Update an organization secret.
     $ {{.HelpName}} --secret.engine native --secret.type org --org MyOrg --name foo --value bar
   4. Update a shared secret.
     $ {{.HelpName}} --secret.engine native --secret.type shared --org MyOrg --team octokitties --name foo --value bar
   5. Update a repository secret with all event types enabled.
     $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --event comment --event deployment --event pull_request --event push --event tag
   6. Update a repository secret with an image whitelist.
     $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --image alpine --image golang:* --image postgres:latest
   7. Update a secret with value from a file.
     $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --value @secret.txt
   8. Update a repository secret with json output.
     $ {{.HelpName}} --secret.engine native --secret.type repo --org MyOrg --repo MyRepo --name foo --value bar --output json
   9. Update a secret or secrets from a file.
     $ {{.HelpName}} --file secret.yml
  10. Update a secret when config or environment variables are set.
     $ {{.HelpName}} --org MyOrg --repo MyRepo --name foo --value bar

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/secret/update/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to modify a secret.
//

func update(ctx context.Context, c *cli.Command) error {
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
		Action:      internal.ActionUpdate,
		Engine:      c.String(internal.FlagSecretEngine),
		Type:        c.String(internal.FlagSecretType),
		Org:         c.String(internal.FlagOrg),
		Repo:        c.String(internal.FlagRepo),
		Team:        c.String("team"),
		Name:        c.String("name"),
		Value:       c.String("value"),
		Images:      c.StringSlice("image"),
		AllowEvents: c.StringSlice("event"),
		File:        c.String("file"),
		Output:      c.String(internal.FlagOutput),
		Color:       output.ColorOptionsFromCLIContext(c),
	}

	// check if allow_command and allow_substitution are provided
	// if they are not, server will not update the fields
	if slices.Contains(c.FlagNames(), internal.FlagSecretCommands) {
		val := c.Bool(internal.FlagSecretCommands)
		s.AllowCommand = &val
	}

	if slices.Contains(c.FlagNames(), internal.FlagSecretSubstitution) {
		val := c.Bool(internal.FlagSecretSubstitution)
		s.AllowSubstitution = &val
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
