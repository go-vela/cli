// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"encoding/json"
	"fmt"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v2"
)

// ViewCmd defines the command for viewing a deployment.
var ViewCmd = cli.Command{
	Name:        "deployment",
	Description: "Use this command to view a deployment.",
	Usage:       "View details of the provided deployment",
	Action:      view,
	Before:      validate,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		&cli.StringFlag{
			Name:    "org",
			Usage:   "Provide the organization for the repository",
			EnvVars: []string{"VELA_ORG"},
		},
		&cli.StringFlag{
			Name:    "repo",
			Usage:   "Provide the repository contained within the organization",
			EnvVars: []string{"VELA_REPO"},
		},
		&cli.IntFlag{
			Name:    "deployment",
			Aliases: []string{"number", "d"},
			Usage:   "Provide the deployment number",
			EnvVars: []string{"VELA_DEPLOYMENT"},
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
 1. View deployment details for a repository.
    $ {{.HelpName}} --org github --repo octocat --deployment 1
 2. View deployment details for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --deployment --output json
 3. View deployment details for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --deployment 1
`, cli.CommandHelpTemplate),
}

// helper function to execute vela info deployment cli command
func view(c *cli.Context) error {
	if c.Int("deployment") == 0 {
		return util.InvalidCommand("deployment")
	}

	// get org, repo and number information from cmd flags
	org, repo, number := c.String("org"), c.String("repo"), c.Int("deployment")

	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	deployment, _, err := client.Deployment.Get(org, repo, number)
	if err != nil {
		return err
	}

	switch c.String("output") {
	case "json":
		output, err := json.MarshalIndent(deployment, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
	default:
		// default output should contain all resources fields
		output, err := yaml.Marshal(deployment)
		if err != nil {
			return err
		}

		fmt.Println(string(output))
	}

	return nil
}
