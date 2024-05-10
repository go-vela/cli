// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
)

// Get captures a list of secrets based on the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	logrus.Debug("executing get for secret configuration")

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
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ListOptions
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	logrus.Tracef("capturing secrets for %s/%s/%s/%s", c.Engine, c.Type, c.Org, name)

	// send API call to capture a list of secrets
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#SecretService.GetAll
	secrets, _, err := client.Secret.GetAll(c.Engine, c.Type, c.Org, name, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the secrets in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(secrets)
	case output.DriverJSON:
		// output the secrets in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(secrets)
	case output.DriverSpew:
		// output the secrets in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(secrets)
	case "wide":
		// output the secrets in wide table format
		return wideTable(secrets)
	case output.DriverYAML:
		// output the secrets in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(secrets, c.Color)
	default:
		// output the secrets in table format
		return table(secrets)
	}
}
