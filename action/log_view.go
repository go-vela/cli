// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/log"
	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/client"

	"github.com/urfave/cli/v2"
)

// LogView defines the command for inspecting a log.
var LogView = &cli.Command{
	Name:        "log",
	Description: "Use this command to view a log.",
	Usage:       "View details of the provided log",
	Action:      logView,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "LOG_ORG"},
			Name:    internal.FlagOrg,
			Aliases: []string{"o"},
			Usage:   "provide the organization for the log",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "LOG_REPO"},
			Name:    internal.FlagRepo,
			Aliases: []string{"r"},
			Usage:   "provide the repository for the log",
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD", "LOG_BUILD"},
			Name:    internal.FlagBuild,
			Aliases: []string{"b"},
			Usage:   "provide the build for the log",
		},

		// Service Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_SERVICE", "LOG_SERVICE"},
			Name:    internal.FlagService,
			Usage:   "provide the service for the log",
		},

		// Step Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_STEP", "LOG_STEP"},
			Name:    internal.FlagStep,
			Usage:   "provide the step for the log",
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
  1. View logs for a build.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1
  2. View logs for a service.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --service 1
  3. View logs for a step.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --step 1
  4. View logs for a build with yaml output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --output yaml
  5. View logs for a build with json output.
    $ {{.HelpName}} --org MyOrg --repo MyRepo --build 1 --output json
  6. View logs for a build when config or environment variables are set.
    $ {{.HelpName}} --build 1

DOCUMENTATION:

  https://go-vela.github.io/docs/reference/cli/log/view/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// inspect a log.
func logView(c *cli.Context) error {
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

	// create the log configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/log?tab=doc#Config
	l := &log.Config{
		Action:  viewAction,
		Org:     c.String(internal.FlagOrg),
		Repo:    c.String(internal.FlagRepo),
		Build:   c.Int(internal.FlagBuild),
		Service: c.Int(internal.FlagService),
		Step:    c.Int(internal.FlagStep),
		Output:  c.String(internal.FlagOutput),
	}

	// validate log configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/log?tab=doc#Config.Validate
	err = l.Validate()
	if err != nil {
		return err
	}

	// check if log service is provided
	if l.Service > 0 {
		// execute the view service call for the log configuration
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/log?tab=doc#Config.ViewService
		return l.ViewService(client)
	}

	// check if log step is provided
	if l.Step > 0 {
		// execute the view step call for the log configuration
		//
		// https://pkg.go.dev/github.com/go-vela/cli/action/log?tab=doc#Config.ViewStep
		return l.ViewStep(client)
	}

	// execute the get call for the log configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/log?tab=doc#Config.Get
	return l.Get(client)
}
