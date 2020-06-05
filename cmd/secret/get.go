// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"

	"github.com/gosuri/uitable"

	"github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v2"
)

// GetCmd defines the command for retrieving secrets from a repository.
var GetCmd = cli.Command{
	Name:        "secret",
	Aliases:     []string{"secrets"},
	Description: "Use this command to get a list of secrets.",
	Usage:       "Display a list of secrets",
	Action:      get,
	Before:      loadGlobal,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		&cli.StringFlag{
			Name:    "engine",
			Usage:   "Provide the engine for where the secret to be stored",
			EnvVars: []string{"VELA_SECRET_ENGINE", "SECRET_ENGINE"},
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

		// optional flags that can be supplied to a command
		&cli.IntFlag{
			Name:    "page",
			Aliases: []string{"p"},
			Usage:   "Print the out the builds a specific page",
			Value:   1,
		},
		&cli.IntFlag{
			Name:    "per-page",
			Aliases: []string{"pp"},
			Usage:   "Expand the number of items contained within page",
			Value:   10,
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Print the output in a yaml or json format",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Get repository secrets.
    $ {{.HelpName}} --engine native --type repo --org MyOrg --repo HelloWorld
 2. Get organization secrets.
    $ {{.HelpName}} --engine native --type org --org MyOrg --repo '*'
 3. Get shared secrets.
    $ {{.HelpName}} --engine native --type shared --org MyOrg --team octokitties
 4. Get secrets for a repository with wide view output.
    $ {{.HelpName}} --output wide --engine native --type repo --org MyOrg --repo HelloWorld
 5. Get secrets for a repository with yaml output.
    $ {{.HelpName}} --output yaml --engine native --type repo --org MyOrg --repo HelloWorld
 6. Get secrets for a repository with json output.
    $ {{.HelpName}} --output json --engine native --type repo --org MyOrg --repo HelloWorld
 7. Get repository secrets with default native engine or when engine and type environment variables are set.
    $ {{.HelpName}} --org MyOrg --repo HelloWorld
`, cli.CommandHelpTemplate),
}

// helper function to execute vela get secrets cli command
func get(c *cli.Context) error {
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

	engine := c.String("engine")
	sType := c.String("type")
	org := c.String("org")
	repo := c.String("repo")

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

	// set the page options based on user input
	opts := &vela.ListOptions{
		Page:    c.Int("page"),
		PerPage: c.Int("per-page"),
	}

	secrets, _, err := client.Secret.GetAll(engine, sType, org, tName, opts)
	if err != nil {
		return err
	}

	switch c.String("output") {
	case "json":
		output, err := json.MarshalIndent(secrets, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))

	case "yaml":
		output, err := yaml.Marshal(secrets)
		if err != nil {
			return err
		}

		fmt.Println(string(output))

	case "wide":
		table := uitable.New()
		table.MaxColWidth = 200
		table.Wrap = true
		table.AddRow("NAME", "ORG", "TYPE", "KEY", "EVENTS", "IMAGES")

		for _, s := range *secrets {
			key, err := getKey(&s)
			if err != nil {
				return fmt.Errorf("invalid key in secret %s: %v", s.GetName(), err)
			}

			if s.Images == nil {
				table.AddRow(s.GetName(), s.GetOrg(), s.GetType(), key, strings.Join(s.GetEvents(), ","), strings.Join(s.GetImages(), ","))
			} else {
				table.AddRow(s.GetName(), s.GetOrg(), s.GetType(), key, strings.Join(s.GetEvents(), ","), strings.Join(s.GetImages(), ","))
			}
		}

		fmt.Println(table)

	default:
		table := uitable.New()
		table.MaxColWidth = 200
		table.Wrap = true // wrap columns
		table.AddRow("NAME", "ORG", "TYPE", "KEY")

		for _, s := range *secrets {
			key, err := getKey(&s)
			if err != nil {
				return fmt.Errorf("invalid key in secret %s: %v", s.GetName(), err)
			}

			table.AddRow(s.GetName(), s.GetOrg(), s.GetType(), key)
		}

		fmt.Println(table)
	}

	return nil
}

// helper function to create a key field from a secret
func getKey(s *library.Secret) (string, error) {
	switch s.GetType() {
	case constants.SecretShared:
		return fmt.Sprintf("%s/%s/%s", s.GetOrg(), s.GetTeam(), s.GetName()), nil
	case constants.SecretOrg:
		return fmt.Sprintf("%s/%s", s.GetOrg(), s.GetName()), nil
	case constants.SecretRepo:
		return fmt.Sprintf("%s/%s/%s", s.GetOrg(), s.GetRepo(), s.GetName()), nil
	}

	return "", fmt.Errorf("invalid secret type")
}
