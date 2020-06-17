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
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	// send API call to capture a list of steps
	steps, _, err := client.Step.GetAll(c.Org, c.Repo, c.Build, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the steps in JSON format
		err := output.JSON(steps)
		if err != nil {
			return err
		}
	case "wide":
		// output the steps in wide table format
		err := wideTable(steps)
		if err != nil {
			return err
		}
	case "yaml":
		// output the steps in YAML format
		err := output.YAML(steps)
		if err != nil {
			return err
		}
	default:
		// output the steps in table format
		err := table(steps)
		if err != nil {
			return err
		}
	}

	return nil
}
