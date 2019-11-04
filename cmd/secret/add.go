// Copyright (c) 2019 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/go-vela/cli/util"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"

	"github.com/urfave/cli"
)

// document is a struct that does a secret yaml document used when
// creating secrets from a file
type document struct {
	Metadata struct {
		APIVersion string `yaml:"api_version"`
		Engine     string `yaml:"engine"`
	} `yaml:"metadata"`
	Secrets []library.Secret `yaml:"secrets"`
}

// AddCmd defines the command to add a secret for a repository.
var AddCmd = cli.Command{
	Name:        "secret",
	Description: "Use this command to add a secret.",
	Usage:       "Add a secret",
	Action:      add,
	Flags: []cli.Flag{

		// required flags to be supplied to a command
		cli.StringFlag{
			Name:   "engine",
			Usage:  "Provide the engine for where the secret to be stored",
			EnvVar: "SECRET_ENGINE",
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
		cli.StringFlag{
			Name:   "value",
			Usage:  "Provide the value of the secret",
			EnvVar: "SECRET_VALUE",
		},

		// optional flags that can be supplied to a command
		cli.StringSliceFlag{
			Name:   "image",
			Usage:  "Secret limited to these images",
			EnvVar: "SECRET_IMAGES",
			Value:  &cli.StringSlice{},
		},
		cli.StringSliceFlag{
			Name:   "event",
			Usage:  "Secret limited to these events",
			EnvVar: "SECRET_EVENTS",
			Value:  &cli.StringSlice{constants.EventPush},
		},
		cli.StringFlag{
			Name:  "filename,f",
			Usage: "Filename to use to add the secret or secrets",
		},
	},
	CustomHelpTemplate: fmt.Sprintf(`%s
EXAMPLES:
 1. Add a secret for a repository with push events.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value bar
 2. Add a secret for a org with push events.
    $ {{.HelpName}} --engine native --type org --org github --repo '*' --name foo --value bar
 3. Add a shared secret for the platform
    $ {{.HelpName}} --engine native --type shared --org github --team octokitties --name foo --value bar
 4. Add a secret for a repository with all event types enabled.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value bar --event push --event pull_request --event tag --event deployment
 5. Add a secret from a file.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value @/path/to/file
 6. Add a native repo secret with an image whitelist.
    $ {{.HelpName}} --engine native --type repo --org github --repo octocat --name foo --value bar --image alpine --image golang
 7. Add a repo secret with default native engine or when engine and type environment variables are set.
	$ {{.HelpName}} --org github --repo octocat --name foo --value bar
 8. Add a secret or secrets from a file
    $ {{.HelpName}} -f secret.yml
`, cli.CommandHelpTemplate),
}

// helper function to execute a add repo cli command
func add(c *cli.Context) error {

	// create a carval client
	client, err := vela.NewClient(c.GlobalString("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.GlobalString("token"))

	switch {

	case len(c.String("filename")) != 0:

		err := procCreateFile(c, client)
		if err != nil {
			return err
		}

	default:
		err := procCreateFlag(c, client)
		if err != nil {
			return err
		}
	}

	return nil
}

// helper function to process user input from CLI flags
func procCreateFlag(c *cli.Context, client *vela.Client) error {

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

	if len(c.String("value")) == 0 {
		return util.InvalidCommand("value")
	}

	engine := c.String("engine")
	sType := c.String("type")
	org := c.String("org")

	// resource to create on server
	images, events := c.StringSlice("image"), c.StringSlice("event")
	req := library.Secret{
		Org:    vela.String(c.String("org")),
		Repo:   vela.String(c.String("repo")),
		Team:   vela.String(c.String("team")),
		Name:   vela.String(c.String("name")),
		Images: &images,
		Events: &events,
		Type:   &sType,
	}

	tName, err := getTypeName(req.GetRepo(), req.GetTeam(), sType)
	if err != nil {
		return err
	}

	// set value from user input or file
	req.Value, err = setValue(c.String("value"))
	if err != nil {
		return err
	}

	secret, _, err := client.Secret.Add(engine, sType, org, tName, &req)
	if err != nil {
		return err
	}

	fmt.Printf("secret \"%s\" was created \n", secret.GetName())

	return nil
}

// helper function to process user input from yaml file
func procCreateFile(c *cli.Context, client *vela.Client) error {

	filename, err := filepath.Abs(c.String("filename"))
	if err != nil {
		return fmt.Errorf("Error getting file %s: %v", c.String("filename"), err)
	}

	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Error reading file  %s: %v", c.String("filename"), err)
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

			secret, _, err := client.Secret.Add(docs.Metadata.Engine, s.GetType(), s.GetOrg(), tName, &s)
			if err != nil {
				return err
			}

			fmt.Printf("secret \"%s\" was created \n", secret.GetName())
		}
	}

	return nil
}
