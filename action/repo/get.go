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
	case "json":
		// output the repositories in JSON format
		err := output.JSON(repos)
		if err != nil {
			return err
		}
	case "wide":
		// TODO: create output.Wide function
		//
		// err := output.Wide(repositories)
		// if err != nil {
		// 	return err
		// }
	case "yaml":
		// output the repositories in YAML format
		err := output.YAML(repos)
		if err != nil {
			return err
		}
	default:
		// output the repositories in default format
		err := output.Default(repos)
		if err != nil {
			return err
		}
	}

	return nil
}
