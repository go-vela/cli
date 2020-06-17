// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/hook"

	"github.com/go-vela/sdk-go/vela"

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
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "provide the organization for the hook",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "HOOK_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "provide the repository for the hook",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "HOOK_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "print the output in default, wide, yaml or json format",
		},

		// Pagination Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_PAGE", "HOOK_PAGE"},
			Name:    "page",
			Aliases: []string{"p"},
			Usage:   "print a specific page of hooks",
			Value:   1,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_PER_PAGE", "HOOK_PER_PAGE"},
			Name:    "per.page",
			Aliases: []string{"pp"},
			Usage:   "number of hooks to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get hooks for a repository.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld
  2. Get hooks for a repository with wide view output.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --output wide
  3. Get hooks for a repository with yaml output.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --output yaml
  4. Get hooks for a repository with json output.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --output json
  5. Get hooks for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/hook/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// capture a list of hooks.
func hookGet(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the hook configuration
	h := &hook.Config{
		Action:  getAction,
		Org:     c.String("org"),
		Repo:    c.String("repo"),
		Page:    c.Int("page"),
		PerPage: c.Int("per.page"),
		Output:  c.String("output"),
	}

	// validate hook configuration
	err = h.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the hook configuration
	return h.Get(client)
}
