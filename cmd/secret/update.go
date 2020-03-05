// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
	"github.com/urfave/cli"
	yaml "gopkg.in/yaml.v2"
)

// UpdateCmd defines the command to update a secret.
var UpdateCmd = cli.Command{
	Name:        "secret",
	Description: "Use this command to update a secret.",
	Usage:       "Update a secret",
	Action:      update,
	Flags: []cli.Flag{

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
			Name:   "value",
			Usage:  "Provide the value of the secret",
			EnvVar: "SECRET_VALUE",
		},
		cli.StringSliceFlag{
			Name:   "image",
			Usage:  "Secret limited to these images",
			EnvVar: "SECRET_IMAGES",
		},
		cli.StringSliceFlag{
			Name:   "event",
			Usage:  "Secret limited to these events",
			EnvVar: "SECRET_EVENTS",
		},
		cli.StringFlag{
			Name:  "filename,f",
			Usage: "Filename to use to create the secret or secrets",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Update a secret value for a repository.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value bar
 2. Update a secret value for an org.
    $ {{.HelpName}} --engine native --type org --org github --repo '*' --name foo --value bar
 3. Update a shared secret value for the platform.
    $ {{.HelpName}} --engine native --type shared --org github --team octokitties --name foo --value bar
 4. Update a secret for a repository with all event types enabled.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --event push --event pull_request --event tag --event deployment
 5. Update a secret from a file.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value @/path/to/file
 6. Update a secret for a repository with an image whitelist.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --image alpine --image golang
 7. Update a repo secret value with default native engine or when engine and type environment variables are set.
	$ {{.HelpName}} --org github --repo octocat --name foo --value bars'
 8. Update with data from a secret file.
	$ {{.HelpName}} -f secret.yml
`, cli.CommandHelpTemplate),
}

// helper function to execute a update repo cli command
func update(c *cli.Context) error {
	// create a carval client
	client, err := vela.NewClient(c.GlobalString("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.GlobalString("token"))

	switch {
	case len(c.String("filename")) != 0:
		err := procUpdateFile(c, client)
		if err != nil {
			return err
		}

	default:
		err := procUpdateFlag(c, client)
		if err != nil {
			return err
		}
	}

	return nil
}

// helper function to process user input from CLI flags
func procUpdateFlag(c *cli.Context, client *vela.Client) error {
	// use global variables if flags aren't provided
	err := loadGlobal(c)
	if err != nil {
		return err
	}

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

	secret.Value, err = setValue(c.String("value"))
	if err != nil {
		return err
	}

	if len(c.StringSlice("image")) > 0 {
		secret.SetImages(c.StringSlice("image"))
	}

	if len(c.StringSlice("event")) > 0 {
		secret.SetEvents(c.StringSlice("event"))
	}

	secret, _, err = client.Secret.Update(engine, sType, org, tName, secret)
	if err != nil {
		return err
	}

	fmt.Printf("secret \"%s\" was updated \n", secret.GetName())

	return nil
}

// helper function to process user input from yaml file
func procUpdateFile(c *cli.Context, client *vela.Client) error {
	filename, err := filepath.Abs(c.String("filename"))
	if err != nil {
		return fmt.Errorf("unable to get file %s: %v", c.String("filename"), err)
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read file  %s: %v", c.String("filename"), err)
	}

	// decode yaml file
	dec := yaml.NewDecoder(bytes.NewReader(yamlFile))

	// create secrets within document
	var docs document
	for dec.Decode(&docs) == nil {
		for _, s := range docs.Secrets {
			tName, err := getTypeName(s.GetRepo(), s.GetTeam(), s.GetType())
			if err != nil {
				return err
			}

			// set value from user input or file

			s.Value, err = setValue(s.GetValue())
			if err != nil {
				return err
			}

			secret, _, err := client.Secret.Update(docs.Metadata.Engine, s.GetType(), s.GetOrg(), tName, &s)
			if err != nil {
				return err
			}

			fmt.Printf("secret \"%s\" was updated \n", secret.GetName())
		}
	}

	return nil
}
