// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"fmt"

	"github.com/go-vela/cli/action/build"

	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// BuildRestart defines the command for restarting a build.
var BuildRestart = &cli.Command{
	Name:        "build",
	Description: "Use this command to restart a build.",
	Usage:       "Re-run the provided build",
	Action:      buildRestart,
	Flags: []cli.Flag{

		// Repo Flags

		&cli.StringFlag{
			EnvVars: []string{"VELA_ORG"},
			Name:    "org",
			Usage:   "Provide the organization for the build",
		},
		&cli.StringFlag{
			EnvVars: []string{"VELA_REPO"},
			Name:    "repo",
			Usage:   "Provide the repository for the build",
		},

		// Build Flags

		&cli.IntFlag{
			EnvVars: []string{"VELA_BUILD"},
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "Provide the build number",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Restart existing build for a repository.
    $ {{.HelpName}} --org github --repo octocat --build-number 1
 2. Restart existing build for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --build-number 1
`, cli.CommandHelpTemplate),
}

// helper function to capture the provided
// input and create the object used to
// restart a build.
func buildRestart(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the build configuration
	b := &build.Build{
		Action: restartAction,
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Number: c.Int("build"),
		Output: c.String("output"),
	}

	// validate build configuration
	err = b.Validate()
	if err != nil {
		return err
	}

	// execute the restart call for the build configuration
	return b.Restart(client)
}
