// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"strings"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/constants"
)

// View inspects a secret based on the provided configuration.
func (c *Config) View(client *vela.Client) error {
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

	// send API call to capture a secret
	secret, _, err := client.Secret.Get(c.Engine, c.Type, c.Org, name, c.Name)
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
