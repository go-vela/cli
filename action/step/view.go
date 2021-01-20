// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package step

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"
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
		return output.JSON(step)
	case output.DriverSpew:
		// output the step in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(step)
	case output.DriverYAML:
		// output the step in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(step)
	default:
		// output the step in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(step)
	}
}
