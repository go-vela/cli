// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"encoding/json"
	"fmt"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

// ViewCmd defines the command for viewing a secret.
var ViewCmd = cli.Command{
	Name:        "secret",
	Description: "Use this command to view a secret.",
	Usage:       "View details of the provided secret",
	Action:      view,
	Before:      loadGlobal,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		cli.StringFlag{
			Name:   "engine",
			Usage:  "Provide the engine for where the secret to be stored",
			EnvVar: "VELA_SECRET_ENGINE,SECRET_ENGINE",
			Value:  constants.DriverNative,
		},
		cli.StringFlag{
			Name:   "type",
			Usage:  "Provide the kind of secret to be stored",
			EnvVar: "SECRET_TYPE",
			Value:  constants.SecretRepo,
		},
		cli.StringFlag{
			Name:   "org",
			Usage:  "Provide the organization for the repository",
			EnvVar: "SECRET_ORG",
		},
		cli.StringFlag{
			Name:   "repo",
			Usage:  "Provide the repository contained with the organization",
			EnvVar: "SECRET_REPO",
		},
		cli.StringFlag{
			Name:   "team",
			Usage:  "Provide the team contained with the organization",
			EnvVar: "SECRET_TEAM",
		},
		cli.StringFlag{
			Name:   "name",
			Usage:  "Provide the name of the secret",
			EnvVar: "SECRET_NAME",
		},

		// optional flags that can be supplied to a command
		cli.StringFlag{
			Name:  "output,o",
			Usage: "Print the output in json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. View repository secret details.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo
 2. View organization secret details.
    $ {{.HelpName}} --engine native --type org --org github --repo '*' --name foo
 3. View shared secret details.
    $ {{.HelpName}} --engine native --type shared --org github --team octokitties --name foo
 4. View secret details for a repository with json output.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --output json
 5. View secret details with default native engine or when engine and type environment variables are set.
    $ {{.HelpName}} --org github --repo octocat --name foo
`, cli.CommandHelpTemplate),
}

// helper function to execute logs cli command
func view(c *cli.Context) error {

	// ensures engine, type, and org are set
	err := validateCmd(c)
	if err != nil {
		return err
	}

	if len(c.String("name")) == 0 {
		return util.InvalidCommand("name")
	}

	tName, err := getTypeName(c.String("repo"), c.String("team"), c.String("type"))
	if err != nil {
		return err
	}

	engine := c.String("engine")
	sType := c.String("type")
	org := c.String("org")
	name := c.String("name")

	// create a vela client
	client, err := vela.NewClient(c.GlobalString("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.GlobalString("token"))

	secret, _, err := client.Secret.Get(engine, sType, org, tName, name)
	if err != nil {
		return err
	}

	switch c.String("output") {
	case "json":
		output, err := json.MarshalIndent(secret, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
	default:
		// default output should contain all resources fields
		output, err := yaml.Marshal(secret)
		if err != nil {
			return err
		}

		fmt.Println(string(output))
	}

	return nil
}
