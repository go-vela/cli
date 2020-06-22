// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// Repair recreates a repository webhook based off the provided configuration.
func (c *Config) Repair(client *vela.Client) error {
	// send API call to repair a repository
	msg, _, err := client.Repo.Repair(c.Org, c.Name)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the message in dump format
		err := output.JSON(msg)
		if err != nil {
			return err
		}
	case output.DriverJSON:
		// output the message in JSON format
		err := output.JSON(msg)
		if err != nil {
			return err
		}
	case output.DriverSpew:
		// output the message in spew format
		err := output.Spew(msg)
		if err != nil {
			return err
		}
	case output.DriverYAML:
		// output the message in YAML format
		err := output.YAML(msg)
		if err != nil {
			return err
		}
	default:
		// output the message in default format
		err := output.Default(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
