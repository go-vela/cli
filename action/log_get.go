// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/log"

	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// LogGet defines the command for capturing a list of build logs.
var LogGet = &cli.Command{
	Name:        "log",
	Aliases:     []string{"logs"},
	Description: "Use this command to get a list of build logs.",
	Usage:       "Display a list of build logs",
	Action:      logGet,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG", "LOG_ORG"},
			Name:    "org",
			Aliases: []string{"o"},
			Usage:   "provide the organization for the log",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO", "LOG_REPO"},
			Name:    "repo",
			Aliases: []string{"r"},
			Usage:   "provide the repository for the log",
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD", "LOG_BUILD"},
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "provide the build for the log",
		},

		// Output Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_OUTPUT", "LOG_OUTPUT"},
			Name:    "output",
			Aliases: []string{"op"},
			Usage:   "print the output in default, yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
  1. Get logs for a build.
    $ {{.HelpName}} --org github --repo octocat --build 1
  2. Get logs for a build with yaml output.
    $ {{.HelpName}} --org github --repo octocat --build 1 --output yaml
  3. Get logs for a build with json output.
    $ {{.HelpName}} --org github --repo octocat --build 1 --output json
  4. Get logs for a build when org and repo config or environment variables are set.
    $ {{.HelpName}} --build 1

DOCUMENTATION:

  https://go-vela.github.io/docs/cli/log/get/
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// capture a list of build logs.
func logGet(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the log configuration
	s := &log.Config{
		Action: getAction,
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Build:  c.Int("build"),
		Output: c.String("output"),
	}

	// validate log configuration
	err = s.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the log configuration
	return s.Get(client)
}
