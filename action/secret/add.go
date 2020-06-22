// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"

	yaml "gopkg.in/yaml.v2"
)

// Add creates a secret based on the provided configuration.
func (c *Config) Add(client *vela.Client) error {
	// check if the secret type is org
	if strings.EqualFold(c.Type, constants.SecretOrg) {
		// set default for the secret repo
		c.Repo = "*"
	}

	// provide the repo name for the secret
	name := c.Repo

	// check if secret type is shared
	if strings.EqualFold(c.Type, constants.SecretShared) {
		// provide the team name for the secret
		name = c.Team
	}

	// create the secret object
	s := &library.Secret{
		Type:         &c.Type,
		Org:          &c.Org,
		Repo:         &c.Repo,
		Team:         &c.Team,
		Name:         &c.Name,
		Value:        &c.Value,
		Images:       &c.Images,
		Events:       &c.Events,
		AllowCommand: &c.AllowCommand,
	}

	// send API call to add a secret
	secret, _, err := client.Secret.Add(c.Engine, c.Type, c.Org, name, s)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the secret in dump format
		err := output.Dump(secret)
		if err != nil {
			return err
		}
	case output.DriverJSON:
		// output the secret in JSON format
		err := output.JSON(secret)
		if err != nil {
			return err
		}
	case output.DriverSpew:
		// output the secret in spew format
		err := output.Spew(secret)
		if err != nil {
			return err
		}
	case output.DriverYAML:
		// output the secret in YAML format
		err := output.YAML(secret)
		if err != nil {
			return err
		}
	default:
		// output the secret in default format
		err := output.Default(secret)
		if err != nil {
			return err
		}
	}

	return nil
}

// AddFromFile creates a secret from a file based on the provided configuration.
func (c *Config) AddFromFile(client *vela.Client) error {
	// capture absolute path to secret file
	path, err := filepath.Abs(c.File)
	if err != nil {
		return err
	}

	// read contents of secret file
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// create a new decoder from the secret file contents
	input := yaml.NewDecoder(bytes.NewReader(contents))

	// create object to store secret file configuration
	f := new(ConfigFile)

	// iterate through all secret file configurations
	for input.Decode(f) == nil {
		// iterate through all secrets from the file configuration
		for _, s := range f.Secrets {
			// create the secret configuration
			s := &Config{
				Action:       "add",
				Engine:       f.Metadata.Engine,
				Type:         s.GetType(),
				Org:          s.GetOrg(),
				Repo:         s.GetRepo(),
				Team:         s.GetTeam(),
				Name:         s.GetName(),
				Value:        s.GetValue(),
				Images:       s.GetImages(),
				Events:       s.GetEvents(),
				AllowCommand: s.GetAllowCommand(),
				Output:       c.Output,
			}

			// validate secret configuration
			err = s.Validate()
			if err != nil {
				return err
			}

			// execute the add call for the secret configuration
			err = s.Add(client)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
