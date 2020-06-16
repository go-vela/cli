// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package hook

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of build hooks based on the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	// set the pagination options for list of hooks
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	// send API call to capture a list of hooks
	hooks, _, err := client.Hook.GetAll(c.Org, c.Repo, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the hooks in JSON format
		err := output.JSON(hooks)
		if err != nil {
			return err
		}
	case "yaml":
		// output the hooks in YAML format
		err := output.YAML(hooks)
		if err != nil {
			return err
		}
	default:
		// output the hooks in default format
		err := output.Default(hooks)
		if err != nil {
			return err
		}
	}

	return nil
}
