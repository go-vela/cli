// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package step

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// View inspects a step based off the provided configuration.
func (c *Config) View(client *vela.Client) error {
	// send API call to capture a step
	step, _, err := client.Step.Get(c.Org, c.Repo, c.Build, c.Number)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the step in JSON format
		err := output.JSON(step)
		if err != nil {
			return err
		}
	case "yaml":
		// output the step in YAML format
		err := output.YAML(step)
		if err != nil {
			return err
		}
	default:
		// output the step in default format
		err := output.Default(step)
		if err != nil {
			return err
		}
	}

	return nil
}
