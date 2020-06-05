// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"fmt"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// RestartCmd defines the command for restarting a specific build.
var RestartCmd = cli.Command{
	Name:        "build",
	Description: "Use this command to restart a build.",
	Usage:       "Re-run the provided build",
	Action:      restart,
	Before:      validate,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		&cli.StringFlag{
			Name:    "org",
			Usage:   "Provide the organization for the repository",
			EnvVars: []string{"BUILD_ORG"},
		},
		&cli.StringFlag{
			Name:    "repo",
			Usage:   "Provide the repository contained within the organization",
			EnvVars: []string{"BUILD_REPO"},
		},
		&cli.IntFlag{
			Name:    "build-number",
			Aliases: []string{"build", "b"},
			Usage:   "Provide the build number",
			EnvVars: []string{"BUILD_NUMBER"},
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Restart existing build for a repository.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --build-number 1
 2. Restart existing build for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --build-number 1
`, cli.CommandHelpTemplate),
}

// helper function to execute vela restart build cli command
func restart(c *cli.Context) error {
	if c.Int("build-number") == 0 {
		return util.InvalidCommand("build-number")
	}

	// get org, repo and number information from cmd flags
	org, repo, number := c.String("org"), c.String("repo"), c.Int("build-number")

	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	build, _, err := client.Build.Restart(org, repo, number)
	if err != nil {
		return err
	}

	fmt.Printf("New build \"%s/%s#%d\" was restarted from \"%s/%s#%d\" \n", org, repo, build.GetNumber(), org, repo, number)

	return nil
}
