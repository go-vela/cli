// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of build logs based on the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	// send API call to capture a list of build logs
	logs, _, err := client.Build.GetLogs(c.Org, c.Repo, c.Build)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the build logs in JSON format
		err := output.JSON(logs)
		if err != nil {
			return err
		}
	case "yaml":
		// output the build logs in YAML format
		err := output.YAML(logs)
		if err != nil {
			return err
		}
	default:
		// output the build logs in default format
		err := output.Default(logs)
		if err != nil {
			return err
		}
	}

	return nil
}
