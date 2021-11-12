// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"fmt"

	"github.com/go-vela/cli/action"
	"github.com/go-vela/cli/action/log"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// CommandGet defines the command for capturing a list of build logs.
var CommandGet = &cli.Command{
	Name:        "log",
	Aliases:     []string{"logs"},
	Description: "Use this command to get a list of build logs.",
	Usage:       "Display a list of build logs",
	Action:      get,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "LOG_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the log",
			Value:   client.GetCwdOrg(),
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "LOG_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the log",
			Value:   client.GetCwdRepo(),
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD", "LOG_BUILD"},
			Name:    internal.FlagBuild,
			Aliases: []string{"b"},
			Usage:   "provide the build for the log",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "LOG_OUTPUT"},
			Name:    internal.FlagOutput,
			Aliases: []string{"op"},
			Usage:   "format the output in json, spew or yaml",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get logs for a build.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1
  2. Get logs for a build with yaml output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --output yaml
  3. Get logs for a build with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --output json
  4. Get logs for a build when config or environment variables are set.
    $ {{.HelpName}} --build 1

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/log/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided input
// and create the object used to capture a list
// of build logs.
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

	// create the log configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/log?tab=doc#Config
	l := &log.Config{
		Action: internal.ActionGet,
		Org:    c.String(internal.FlagOrg),
		Repo:   c.String(internal.FlagRepo),
		Build:  c.Int(internal.FlagBuild),
		Output: c.String(internal.FlagOutput),
	}

	// validate log configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/log?tab=doc#Config.Validate
	err = l.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the log configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/log?tab=doc#Config.Get
	return l.Get(client)
}
