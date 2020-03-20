// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package service

import (
	"encoding/json"
	"fmt"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

// ViewCmd defines the command for viewing a service.
var ViewCmd = cli.Command{
	Name:        "service",
	Description: "Use this command to view a service.",
	Usage:       "View details of the provided service",
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
			Usage:  "Provide the repository contained with the organization",
			EnvVar: "BUILD_REPO",
		},
		cli.IntFlag{
			Name:   "build-number,build,b",
			Usage:  "Provide the build number",
			EnvVar: "BUILD_NUMBER",
		},
		cli.IntFlag{
			Name:   "service-number,service,s",
			Usage:  "Provide the service number",
			EnvVar: "SERVICE_NUMBER",
		},

		// optional flags that can be supplied to a command
		cli.StringFlag{
			Name:  "output,o",
			Usage: "Print the output in json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Get service details for a repository.
    $ {{.HelpName}} --org github --repo octocat --build-number 1 --service-number 1
 2. Get service details for a repository with json output.
    $ {{.HelpName}} --org github --repo octocat --build-number 1 --service-number 1 --output json
 3. Get service details for a repository when org and repo config or environment variables are set.
    $ {{.HelpName}} --build-number 1 --service-number 1
`, cli.CommandHelpTemplate),
}

// helper function to execute vela info service cli command
func view(c *cli.Context) error {
	if c.Int("service-number") == 0 {
		return util.InvalidCommand("service-number")
	}

	// get org, repo build and service number information from cmd flags
	org, repo := c.String("org"), c.String("repo")
	bNum, sNum := c.Int("build-number"), c.Int("service-number")

	// create a vela client
	client, err := vela.NewClient(c.GlobalString("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.GlobalString("token"))

	svc, _, err := client.Svc.Get(org, repo, bNum, sNum)
	if err != nil {
		return err
	}

	switch c.String("output") {
	case "json":
		output, err := json.MarshalIndent(svc, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
	default:
		// default output should contain all resources fields
		output, err := yaml.Marshal(svc)
		if err != nil {
			return err
		}

		fmt.Println(string(output))
	}

	return nil
}
