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

// Get captures a list of secrets based on the provided configuration.
func (c *Config) Get(client *vela.Client) error {
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

	// set the pagination options for list of secrets
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	// send API call to capture a list of secrets
	secrets, _, err := client.Secret.GetAll(c.Engine, c.Type, c.Org, name, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the secrets in JSON format
		err := output.JSON(secrets)
		if err != nil {
			return err
		}
	case "wide":
		// output the secrets in wide table format
		err := wideTable(secrets)
		if err != nil {
			return err
		}
	case "yaml":
		// output the secrets in YAML format
		err := output.YAML(secrets)
		if err != nil {
			return err
		}
	default:
		// output the secrets in table format
		err := table(secrets)
		if err != nil {
			return err
		}
	}

	return nil
}
