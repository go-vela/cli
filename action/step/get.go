// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package step

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of steps based on the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	// set the pagination options for list of steps
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ListOptions
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	// send API call to capture a list of steps
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#StepService.GetAll
	steps, _, err := client.Step.GetAll(c.Org, c.Repo, c.Build, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the steps in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(steps)
	case output.DriverJSON:
		// output the steps in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(steps)
	case output.DriverSpew:
		// output the steps in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(steps)
	case "wide":
		// output the steps in wide table format
		return wideTable(steps)
	case output.DriverYAML:
		// output the steps in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(steps)
	default:
		// output the steps in table format
		return table(steps)
	}
}
