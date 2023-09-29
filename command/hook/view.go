// SPDX-License-Identifier: Apache-2.0

package hook

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/hook"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// CommandView defines the command for inspecting a hook.
var CommandView = &cli.Command{
	Name:        "hook",
	Description: "Use this command to view a hook.",
	Usage:       "View details of the provided hook",
	Action:      view,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "HOOK_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the hook",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "HOOK_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the hook",
		},

		// Hook Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_HOOK", "HOOK_NUMBER"},
			Name:    "hook",
			Aliases: []string{"number", "hn"},
			Usage:   "provide the number for the hook",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "HOOK_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
			Value:   "yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. View hook details for a repository.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --hook 1
  2. View hook details for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --hook 1 --output json
  3. View hook details for a repository when config or environment variables are set.
    $ {{.HelpName}} --hook 1

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/hook/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to inspect a hook.
func view(c *cli.Context) error {
	// load variables from the config file
	err := action.Load(c)
	if err != nil {
		return err
	}

	// grab first command line argument, if it exists, and set it as resource
	err = internal.ProcessArgs(c, "hook", "int")
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

	// create the hook configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/hook?tab=doc#Config
	h := &hook.Config{
		Action: internal.ActionView,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Number: c.Int("hook"),
		Output: c.String(internal.FlagOutput),
	}

	// validate hook configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/hook?tab=doc#Config.Validate
	err = h.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the hook configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/hook?tab=doc#Config.View
	return h.View(client)
}
