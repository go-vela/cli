// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of deployments based off the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	// set the pagination options for list of deployments
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	// send API call to capture a list of deployments
	deployments, _, err := client.Deployment.GetAll(c.Org, c.Repo, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the deployments in JSON format
		err := output.JSON(deployments)
		if err != nil {
			return err
		}
	case "wide":
		// output the deployments in wide table format
		err := wideTable(deployments)
		if err != nil {
			return err
		}
	case "yaml":
		// output the deployments in YAML format
		err := output.YAML(deployments)
		if err != nil {
			return err
		}
	default:
		// output the deployments in table format
		err := table(deployments)
		if err != nil {
			return err
		}
	}

	return nil
}
