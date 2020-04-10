// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"fmt"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"
	"github.com/urfave/cli/v2"
)

// repairCmd defines the command for generating a pipeline.
var repairCmd = cli.Command{
	Name:        "repair",
	Category:    "Repository Management",
	Aliases:     []string{"re"},
	Description: "Use this command to repair a damaged repository",
	Usage:       "Repair damaged repositories",
	Action:      repair,
	Before:      validateRepair,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		&cli.StringFlag{
			Name:    "org",
			Usage:   "Provide the organization for the repository",
			EnvVars: []string{"REPO_ORG"},
		},
		&cli.StringFlag{
			Name:    "repo",
			Usage:   "Provide the repository contained with the organization",
			EnvVars: []string{"REPO_NAME"},
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Fix a damaged repository by disabling and enabling."
	$ {{.HelpName}}
`, cli.CommandHelpTemplate),
}

// helper function to run the repair process for a repository
func repair(c *cli.Context) error {

	// get org and repo information from cmd flags
	org, repo := c.String("org"), c.String("repo")

	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from context
	client.Authentication.SetTokenAuth(c.String("token"))

	_, _, err = client.Repo.Repair(org, repo)
	if err != nil {
		return err
	}

	fmt.Printf("repo \"%s/%s\" was repaired \n", org, repo)

	return err
}

// helper function to load global configuration if set
// via config or environment and validate the user input in the command
func validateRepair(c *cli.Context) error {

	// load configuration
	if len(c.String("org")) == 0 {
		c.Set("org", c.String("org"))
	}
	if len(c.String("repo")) == 0 {
		c.Set("repo", c.String("repo"))
	}

	// validate the user input in the command
	if len(c.String("org")) == 0 {
		return util.InvalidCommand("org")
	}
	if len(c.String("repo")) == 0 {
		return util.InvalidCommand("repo")
	}

	return nil
}
