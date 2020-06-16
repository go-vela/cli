// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package hook

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// View inspects a hook based off the provided configuration.
func (c *Config) View(client *vela.Client) error {
	// send API call to capture a hook
	hook, _, err := client.Hook.Get(c.Org, c.Repo, c.Number)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case "json":
		// output the hook in JSON format
		err := output.JSON(hook)
		if err != nil {
			return err
		}
	case "yaml":
		// output the hook in YAML format
		err := output.YAML(hook)
		if err != nil {
			return err
		}
	default:
		// output the hook in default format
		err := output.Default(hook)
		if err != nil {
			return err
		}
	}

	return nil
}
