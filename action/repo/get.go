// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of deployments based off the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	// set the pagination options for list of repositories
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	// send API call to capture a list of repositories
	repos, _, err := client.Repo.GetAll(opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the repositories in dump format
		err := output.Dump(repos)
		if err != nil {
			return err
		}
	case output.DriverJSON:
		// output the repositories in JSON format
		err := output.JSON(repos)
		if err != nil {
			return err
		}
	case output.DriverSpew:
		// output the repositories in spew format
		err := output.Spew(repos)
		if err != nil {
			return err
		}
	case "wide":
		// output the repos in wide table format
		err := wideTable(repos)
		if err != nil {
			return err
		}
	case output.DriverYAML:
		// output the repositories in YAML format
		err := output.YAML(repos)
		if err != nil {
			return err
		}
	default:
		// output the repos in table format
		err := table(repos)
		if err != nil {
			return err
		}
	}

	return nil
}
