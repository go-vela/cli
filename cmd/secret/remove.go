// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"fmt"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"

	"github.com/urfave/cli/v2"
)

// RemoveCmd defines the command to remove a repository.
var RemoveCmd = cli.Command{
	Name:        "secret",
	Description: "Use this command to remove a secret.",
	Usage:       "Remove a secret",
	Action:      remove,
	Before:      loadGlobal,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		&cli.StringFlag{
			Name:    "engine",
			Usage:   "Provide the engine for where the secret to be stored",
			EnvVars: []string{"VELA_SECRET_ENGIN", "SECRET_ENGINE"},
			Value:   constants.DriverNative,
		},
		&cli.StringFlag{
			Name:    "type",
			Usage:   "Provide the kind of secret to be stored",
			EnvVars: []string{"SECRET_TYPE"},
			Value:   constants.SecretRepo,
		},
		&cli.StringFlag{
			Name:    "org",
			Usage:   "Provide the organization for the repository",
			EnvVars: []string{"SECRET_ORG"},
		},
		&cli.StringFlag{
			Name:    "repo",
			Usage:   "Provide the repository contained with the organization",
			EnvVars: []string{"SECRET_REPO"},
		},
		&cli.StringFlag{
			Name:    "team",
			Usage:   "Provide the team contained with the organization",
			EnvVars: []string{"SECRET_TEAM"},
		},
		&cli.StringFlag{
			Name:    "name",
			Usage:   "Provide the name of the secret",
			EnvVars: []string{"SECRET_NAME"},
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Remove a secret for a repository.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo
 2. Remove a secret for an org.
    $ {{.HelpName}} --engine native --type org --org github --repo '*' --name foo
 3. Remove a shared secret for the platform.
    $ {{.HelpName}} --engine native --type shared --org github --team octokitties --name foo
 4. Remove a repo secret with default native engine or when engine and type environment variables are set.
    $ {{.HelpName}} --org github --repo octocat --name foo
`, cli.CommandHelpTemplate),
}

// helper function to execute a remove repo cli command
func remove(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// ensures engine, type, and org are set
	err = validateCmd(c)
	if err != nil {
		return err
	}

	if len(c.String("name")) == 0 {
		return util.InvalidCommand("name")
	}

	engine := c.String("engine")
	sType := c.String("type")
	org := c.String("org")
	repo := c.String("repo")
	name := c.String("name")

	// check if the secret provided is an org type
	if sType == constants.SecretOrg {
		// check if the repo was provided
		if len(repo) == 0 {
			// set a default for the repo
			repo = "*"
		}
	}

	tName, err := getTypeName(repo, name, sType)
	if err != nil {
		return err
	}

	_, _, err = client.Secret.Remove(engine, sType, org, tName, name)
	if err != nil {
		return err
	}

	fmt.Printf("secret \"%s\" was removed \n", name)

	return nil
}
