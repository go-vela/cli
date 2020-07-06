// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"strings"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/constants"

	"github.com/sirupsen/logrus"
)

// View inspects a secret based on the provided configuration.
func (c *Config) View(client *vela.Client) error {
	logrus.Debug("executing view for secret configuration")

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

	logrus.Tracef("inspecting secret %s/%s/%s/%s/%s", c.Engine, c.Type, c.Org, name, c.Name)

	// send API call to capture a secret
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#SecretService.Get
	secret, _, err := client.Secret.Get(c.Engine, c.Type, c.Org, name, c.Name)
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
