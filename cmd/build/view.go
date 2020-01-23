// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"encoding/json"
	"fmt"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

// ViewCmd defines the command for viewing a build.
var ViewCmd = cli.Command{
	Name:        "build",
	Description: "Use this command to view a build.",
	Usage:       "View details of the provided build",
	Action:      view,
	Before:      validate,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		cli.StringFlag{
			Name:   "org",
			Usage:  "Provide the organization for the repository",
			EnvVar: "BUILD_ORG",
		},
		cli.StringFlag{
			Name:   "repo",
			Usage:  "Provide the repository contained within the organization",
			EnvVar: "BUILD_REPO",
		},
		cli.IntFlag{
			Name:   "build-number,build,b",
			Usage:  "Provide the build number",
			EnvVar: "BUILD_NUMBER",
		},

		// optional flags that can be supplied to a command
		cli.StringFlag{
			Name:  "output,o",
			Usage: "Print the output in json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. View build details for a repository.
    $ {{.HelpName}} --org github --repo octocat --build-number 1
 2. View build details for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --build-number 1 --output json
 3. View build details for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --build-number 1
`, cli.CommandHelpTemplate),
}

// helper function to execute vela info build cli command
func view(c *cli.Context) error {
	if c.Int("build-number") == 0 {
		return util.InvalidCommand("build-number")
	}

	// get org, repo and number information from cmd flags
	org, repo, number := c.String("org"), c.String("repo"), c.Int("build-number")

	// create a carval client
	client, err := vela.NewClient(c.GlobalString("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.GlobalString("token"))

	build, _, err := client.Build.Get(org, repo, number)
	if err != nil {
		return err
	}

	switch c.String("output") {
	case "json":
		output, err := json.MarshalIndent(build, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
	default:
		// default output should contain all resources fields
		output, err := yaml.Marshal(build)
		if err != nil {
			return err
		}

		fmt.Println(string(output))
	}

	return nil
}
