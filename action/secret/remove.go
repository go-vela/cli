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

// Remove deletes a secret based on the provided configuration.
func (c *Config) Remove(client *vela.Client) error {
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

	// send API call to remove a secret
	msg, _, err := client.Secret.Remove(c.Engine, c.Type, c.Org, name, c.Name)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the message in dump format
		err := output.Dump(msg)
		if err != nil {
			return err
		}
	case output.DriverJSON:
		// output the message in JSON format
		err := output.JSON(msg)
		if err != nil {
			return err
		}
	case output.DriverSpew:
		// output the message in spew format
		err := output.Spew(msg)
		if err != nil {
			return err
		}
	case output.DriverYAML:
		// output the message in YAML format
		err := output.YAML(msg)
		if err != nil {
			return err
		}
	default:
		// output the message in stdout format
		err := output.Stdout(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
