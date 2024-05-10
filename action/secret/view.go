// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v3"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
	pyaml "github.com/go-vela/types/yaml"
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
		return output.YAML(secret, c.Color)
	default:
		return outputDefault(c.Engine, secret)
	}
}

// outputDefault is a helper function to output the
// provided secrets with a copy pipeline secret and
// library details formatted with yaml.
func outputDefault(engine string, s *library.Secret) error {
	// create yaml secret
	secret := &pyaml.Secret{
		Name:   s.GetName(),
		Key:    key(s),
		Engine: engine,
		Type:   s.GetType(),
	}

	// anonymous struct for yaml secret
	_secret := struct {
		Secret []*pyaml.Secret `yaml:"secret"`
	}{
		[]*pyaml.Secret{secret},
	}

	// anonymous struct for library secret
	_s := struct {
		Details *library.Secret `yaml:"details"`
	}{
		s,
	}

	// marshal the input into YAML
	output, err := yaml.Marshal(_secret)
	if err != nil {
		return err
	}

	// marshal the input into YAML
	tmp, err := yaml.Marshal(_s)
	if err != nil {
		return err
	}

	// add a new line between the pipeline and details output
	output = append(output, "\n"...)
	output = append(output, tmp...)

	// ensure we output to stdout
	fmt.Fprintln(os.Stdout, string(output))

	return nil
}
