// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of builds based off the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	// set the pagination options for list of builds
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	// send API call to capture a list of builds
	builds, _, err := client.Build.GetAll(c.Org, c.Repo, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the build in JSON format
		err := output.JSON(builds)
		if err != nil {
			return err
		}
	case "wide":
		// TODO: create output.Wide function
		//
		// err := output.Wide(builds)
		// if err != nil {
		// 	return err
		// }
	case "yaml":
		// output the build in YAML format
		err := output.YAML(builds)
		if err != nil {
			return err
		}
	default:
		// output the build in default format
		err := output.Default(builds)
		if err != nil {
			return err
		}
	}

	return nil
}
