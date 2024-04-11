// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
)

// Remove deletes a secret based on the provided configuration.
func (c *Config) Remove(client *vela.Client) error {
	logrus.Debug("executing remove for secret configuration")

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

	logrus.Tracef("removing secret %s/%s/%s/%s/%s", c.Engine, c.Type, c.Org, name, c.Name)

	// send API call to remove a secret
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#SecretService.Remove
	msg, _, err := client.Secret.Remove(c.Engine, c.Type, c.Org, name, c.Name)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the msg in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(msg)
	case output.DriverJSON:
		// output the msg in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(msg)
	case output.DriverSpew:
		// output the msg in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(msg)
	case output.DriverYAML:
		// output the msg in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(msg)
	default:
		// output the msg in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(*msg)
	}
}
