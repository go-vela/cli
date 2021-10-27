// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/repo"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// CommandGet defines the command for capturing a list of repositories.
var CommandGet = &cli.Command{
	Name:        "repo",
	Aliases:     []string{"repos"},
	Description: "Use this command to get a list of repositories.",
	Usage:       "Display a list of repositories",
	Action:      get,
	Flags: []cli.Flag{

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "REPO_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew, wide or yaml",
		},

		// Pagination Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_PAGE", "REPO_PAGE"},
			Name:    internal.FlagPage,
			Aliases: []string{"p"},
			Usage:   "print a specific page of repositories",
			Value:   1,
		},
		&cli.IntFlag{
			EnvVars: []string{"VELA_PER_PAGE", "REPO_PER_PAGE"},
			Name:    internal.FlagPerPage,
			Aliases: []string{"pp"},
			Usage:   "number of repositories to print per page",
			Value:   10,
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get a list of repositories.
    $ {{.HelpName}}
  2. Get a list of repositories with wide view output.
    $ {{.HelpName}} --output wide
  3. Get a list of repositories with yaml output.
    $ {{.HelpName}} --output yaml
  4. Get a list of repositories with json output.
    $ {{.HelpName}} --output json
  5. Get a list of repositories when config or environment variables are set.
    $ {{.HelpName}}

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/repo/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to capture a list
// of repos.
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

	// create the repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config
	r := &repo.Config{
		Action:  internal.ActionGet,
		Page:    c.Int(internal.FlagPage),
		PerPage: c.Int(internal.FlagPerPage),
		Output:  c.String(internal.FlagOutput),
	}

	// validate repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Validate
	err = r.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the repo configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/repo?tab=doc#Config.Get
	return r.Get(client)
}
