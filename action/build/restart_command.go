// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Restart defines the command for restarting a build.
var Restart = &cli.Command{
	Name:        "build",
	Description: "Use this command to restart a build.",
	Usage:       "Re-run the provided build",
	Action:      restart,
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
