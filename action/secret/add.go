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

	"github.com/sirupsen/logrus"

	yaml "gopkg.in/yaml.v2"
)

// Add creates a secret based on the provided configuration.
func (c *Config) Add(client *vela.Client) error {
	logrus.Debug("executing add for secret configuration")

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

	// set the proper value for the secret
	err := c.setValue()
	if err != nil {
		return err
	}

	// create the secret object
	//
	// https://pkg.go.dev/github.com/go-vela/types/library?tab=doc#Secret
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

	logrus.Tracef("adding secret %s/%s/%s/%s/%s", c.Engine, c.Type, c.Org, name, c.Name)

	// send API call to add a secret
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#SecretService.Add
	secret, _, err := client.Secret.Add(c.Engine, c.Type, c.Org, name, s)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the secret in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(secret)
	case output.DriverJSON:
		// output the secret in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(secret)
	case output.DriverSpew:
		// output the secret in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(secret)
	case output.DriverYAML:
		// output the secret in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(secret)
	default:
		// output the secret in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(secret)
	}
}

// AddFromFile creates a secret from a file based on the provided configuration.
func (c *Config) AddFromFile(client *vela.Client) error {
	logrus.Debug("executing add from file for secret configuration")

	// capture absolute path to secret file
	path, err := filepath.Abs(c.File)
	if err != nil {
		return err
	}

	logrus.Tracef("reading secret contents from %s", path)

	// read contents of secret file
	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// create a new decoder from the secret file contents
	//
	// https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc#NewDecoder
	input := yaml.NewDecoder(bytes.NewReader(contents))

	// create object to store secret file configuration
	//
	// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#ConfigFile
	f := new(ConfigFile)

	// iterate through all secret file configurations
	//
	// https://pkg.go.dev/gopkg.in/yaml.v2?tab=doc#Decoder.Decode
	for input.Decode(f) == nil {
		// iterate through all secrets from the file configuration
		for _, s := range f.Secrets {
			// create the secret configuration
			//
			// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config
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
			//
			// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.Validate
			err = s.Validate()
			if err != nil {
				return err
			}

			// execute the add call for the secret configuration
			//
			// https://pkg.go.dev/github.com/go-vela/cli/action/secret?tab=doc#Config.Add
			err = s.Add(client)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
