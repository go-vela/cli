// SPDX-License-Identifier: Apache-2.0

package step

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// View inspects a step based on the provided configuration.
func (c *Config) View(client *vela.Client) error {
	logrus.Debug("executing view for step configuration")

	logrus.Tracef("inspecting step %s/%s/%d/%d", c.Org, c.Repo, c.Build, c.Number)

	// send API call to capture a step
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#StepService.Get
	step, _, err := client.Step.Get(c.Org, c.Repo, c.Build, c.Number)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the step in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(step)
	case output.DriverJSON:
		// output the step in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(step, c.Color)
	case output.DriverSpew:
		// output the step in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(step)
	case output.DriverYAML:
		// output the step in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(step, c.Color)
	default:
		// output the step in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(step)
	}
}
