// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// ViewService inspects a service log based on the provided configuration.
func (c *Config) ViewService(client *vela.Client) error {
	// send API call to capture a service log
	log, _, err := client.Log.GetService(c.Org, c.Repo, c.Build, c.Service)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the service log in JSON format
		err := output.JSON(log)
		if err != nil {
			return err
		}
	case "yaml":
		// output the service log in YAML format
		err := output.YAML(log)
		if err != nil {
			return err
		}
	default:
		// output the service log in default format
		err := output.Default(log)
		if err != nil {
			return err
		}
	}

	return nil
}

// ViewStep inspects a service log based on the provided configuration.
func (c *Config) ViewStep(client *vela.Client) error {
	// send API call to capture a step log
	log, _, err := client.Log.GetStep(c.Org, c.Repo, c.Build, c.Step)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the step log in JSON format
		err := output.JSON(log)
		if err != nil {
			return err
		}
	case "yaml":
		// output the step log in YAML format
		err := output.YAML(log)
		if err != nil {
			return err
		}
	default:
		// output the step log in default format
		err := output.Default(log)
		if err != nil {
			return err
		}
	}

	return nil
}
