// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code among actions
package build

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Restart restarts a build based off the provided configuration.
func (c *Config) Restart(client *vela.Client) error {
	logrus.Debug("executing restart for build configuration")

	logrus.Tracef("restarting build %s/%s/%d", c.Org, c.Repo, c.Number)

	// send API call to restart a build
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#BuildService.Restart
	build, _, err := client.Build.Restart(c.Org, c.Repo, c.Number)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the build in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(build)
	case output.DriverJSON:
		// output the build in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(build, c.Color)
	case output.DriverSpew:
		// output the build in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(build)
	case output.DriverYAML:
		// output the build in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(build, c.Color)
	default:
		// output the build in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(build)
	}
}
