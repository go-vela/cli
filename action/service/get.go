// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package service

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of services based off the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	// set the pagination options for list of services
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	// send API call to capture a list of services
	services, _, err := client.Svc.GetAll(c.Org, c.Repo, c.Build, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the service in JSON format
		err := output.JSON(services)
		if err != nil {
			return err
		}
	case "wide":
		// TODO: create output.Wide function
		//
		// err := output.Wide(services)
		// if err != nil {
		// 	return err
		// }
	case "yaml":
		// output the service in YAML format
		err := output.YAML(services)
		if err != nil {
			return err
		}
	default:
		// output the service in default format
		err := output.Default(services)
		if err != nil {
			return err
		}
	}

	return nil
}
