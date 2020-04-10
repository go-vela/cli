// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"fmt"

	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// ViewCmd defines the command for viewing the logs from a build or step.
var ViewCmd = cli.Command{
	Name:        "log",
	Aliases:     []string{"logs"},
	Description: "Use this command to capture the logs from a build or step.",
	Usage:       "View logs from the provided build or step",
	Action:      view,
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
			Usage:   "Provide the repository contained with the organization",
			EnvVars: []string{"BUILD_REPO"},
		},
		&cli.IntFlag{
			Name:    "build-number",
			Aliases: []string{"build", "b"},
			Usage:   "Print the output in wide, yaml or json format",
			EnvVars: []string{"BUILD_NUMBER"},
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. View build logs for a repository.
    $ {{.HelpName}} --org github --repo octocat --build-number 1
 2. View build logs for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --build-number 1
`, cli.CommandHelpTemplate),
}

// helper function to execute logs cli command
func view(c *cli.Context) error {

	// get org, repo and number information from cmd flags
	org, repo, number := c.String("org"), c.String("repo"), c.Int("build-number")

	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// Get the build you just created
	build, _, err := client.Build.Get(org, repo, number)
	if err != nil {
		return err
	}

	// Get the build you just created
	logs, _, err := client.Build.GetLogs(org, repo, build.GetNumber())
	if err != nil {
		return err
	}

	// print logs for all steps
	for _, log := range *logs {
		fmt.Printf("%s \n", log.GetData())
	}

	return nil
}
