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
	case output.DriverDump:
		// output the builds in dump format
		err := output.Dump(builds)
		if err != nil {
			return err
		}
	case output.DriverJSON:
		// output the builds in JSON format
		err := output.JSON(builds)
		if err != nil {
			return err
		}
	case output.DriverSpew:
		// output the builds in spew format
		err := output.Spew(builds)
		if err != nil {
			return err
		}
	case "wide":
		// output the builds in wide table format
		err := wideTable(builds)
		if err != nil {
			return err
		}
	case output.DriverYAML:
		// output the builds in YAML format
		err := output.YAML(builds)
		if err != nil {
			return err
		}
	default:
		// output the builds in table format
		err := table(builds)
		if err != nil {
			return err
		}
	}

	return nil
}
