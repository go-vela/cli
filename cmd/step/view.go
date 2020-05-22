// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package step

import (
	"encoding/json"
	"fmt"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v2"
)

// ViewCmd defines the command for viewing a step.
var ViewCmd = cli.Command{
	Name:        "step",
	Description: "Use this command to view a step.",
	Usage:       "View details of the provided step",
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
			Usage:   "Provide the build number",
			EnvVars: []string{"BUILD_NUMBER"},
		},
		&cli.IntFlag{
			Name:    "step-number",
			Aliases: []string{"step", "s"},
			Usage:   "Provide the step number",
			EnvVars: []string{"STEP_NUMBER"},
		},

		// optional flags that can be supplied to a command
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Print the output in json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Get step details for a repository.
    $ {{.HelpName}} --org github --repo octocat --build-number 1 --step-number 1
 2. Get step details for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --build-number 1 --step-number 1 --output json
 3. Get step details for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --build-number 1 --step-number 1
`, cli.CommandHelpTemplate),
}

// helper function to execute vela info step cli command
func view(c *cli.Context) error {
	if c.Int("step-number") == 0 {
		return util.InvalidCommand("step-number")
	}

	// get org, repo build and step number information from cmd flags
	org, repo := c.String("org"), c.String("repo")
	bNum, sNum := c.Int("build-number"), c.Int("step-number")

	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	build, _, err := client.Step.Get(org, repo, bNum, sNum)
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
