// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package service

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// View inspects a service based on the provided configuration.
func (c *Config) View(client *vela.Client) error {
	// send API call to capture a service
	service, _, err := client.Svc.Get(c.Org, c.Repo, c.Build, c.Number)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the service in dump format
		err := output.Dump(service)
		if err != nil {
			return err
		}
	case output.DriverJSON:
		// output the service in JSON format
		err := output.JSON(service)
		if err != nil {
			return err
		}
	case output.DriverSpew:
		// output the service in spew format
		err := output.Spew(service)
		if err != nil {
			return err
		}
	case output.DriverYAML:
		// output the service in YAML format
		err := output.YAML(service)
		if err != nil {
			return err
		}
	default:
		// output the service in stdout format
		err := output.Stdout(service)
		if err != nil {
			return err
		}
	}

	return nil
}
