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
	"github.com/urfave/cli/v2"
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

		// optional flags that can be supplied to a command
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Print the output in json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. View repository secret details.
    $ {{.HelpName}} --engine native --type repo --org MyOrg --repo HelloWorld --name foo
 2. View organization secret details.
    $ {{.HelpName}} --engine native --type org --org MyOrg --repo '*' --name foo
 3. View shared secret details.
    $ {{.HelpName}} --engine native --type shared --org MyOrg --team octokitties --name foo
 4. View secret details for a repository with json output.
    $ {{.HelpName}} --engine native --type repo --org MyOrg --repo HelloWorld --name foo --output json
 5. View secret details with default native engine or when engine and type environment variables are set.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld --name foo
`, cli.CommandHelpTemplate),
}

// helper function to execute logs cli command
func view(c *cli.Context) error {
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

	tName, err := getTypeName(repo, c.String("team"), sType)
	if err != nil {
		return err
	}

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
