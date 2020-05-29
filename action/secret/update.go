// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"strings"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
)

// Update modifies a secret based off the provided configuration.
func (c *Config) Update(client *vela.Client) error {
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
		Type:   &c.Type,
		Org:    &c.Org,
		Repo:   &c.Repo,
		Team:   &c.Team,
		Name:   &c.Name,
		Images: &c.Images,
		Events: &c.Events,
	}

	// send API call to update a secret
	secret, _, err := client.Secret.Update(c.Engine, c.Type, c.Org, name, s)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the secret in JSON format
		err := output.JSON(secret)
		if err != nil {
			return err
		}
	case "yaml":
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
