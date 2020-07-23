// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/hook"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// HookGet defines the command for capturing a list of hooks.
var HookGet = &cli.Command{
	Name:        "hook",
	Aliases:     []string{"hooks"},
	Description: "Use this command to get a list of hooks.",
	Usage:       "Display a list of hooks",
	Action:      hookGet,
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

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "HOOK_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},

		// Pagination Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_PAGE", "HOOK_PAGE"},
			Name:    internal.FlagPage,
			Aliases: []string{"p"},
			Usage:   "print a specific page of hooks",
			Value:   1,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_PER_PAGE", "HOOK_PER_PAGE"},
			Name:    internal.FlagPerPage,
			Aliases: []string{"pp"},
			Usage:   "number of hooks to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get hooks for a repository.
    $ {{.HelpName}} --org MyOrg --repo MyRepo
  2. Get hooks for a repository with wide view output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --output wide
  3. Get hooks for a repository with yaml output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --output yaml
  4. Get hooks for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --output json
  5. Get hooks for a repository when config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/hook/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// capture a list of hooks.
func hookGet(c *cli.Context) error {
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
		Action:  getAction,
		Org:     c.String(internal.FlagOrg),
		Repo:    c.String(internal.FlagRepo),
		Page:    c.Int(internal.FlagPage),
		PerPage: c.Int(internal.FlagPerPage),
		Output:  c.String(internal.FlagOutput),
	}

	// validate hook configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/hook?tab=doc#Config.Validate
	err = h.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the hook configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/hook?tab=doc#Config.Get
	return h.Get(client)
}
