// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/library"
)

// Add creates a deployment based off the provided configuration.
func (c *Config) Add(client *vela.Client) error {
	// create the deployment object
	d := &library.Deployment{
		Ref:         &c.Ref,
		Task:        &c.Task,
		Target:      &c.Target,
		Description: &c.Description,
	}

	// send API call to add a deployment
	deployment, _, err := client.Deployment.Add(c.Org, c.Repo, d)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the deployment in dump format
		err := output.Dump(deployment)
		if err != nil {
			return err
		}
	case output.DriverJSON:
		// output the deployment in JSON format
		err := output.JSON(deployment)
		if err != nil {
			return err
		}
	case output.DriverSpew:
		// output the deployment in spew format
		err := output.Spew(deployment)
		if err != nil {
			return err
		}
	case output.DriverYAML:
		// output the deployment in YAML format
		err := output.YAML(deployment)
		if err != nil {
			return err
		}
	default:
		// output the deployment in default format
		err := output.Default(deployment)
		if err != nil {
			return err
		}
	}

	return nil
}
