// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"
)

// View inspects a repository based off the provided configuration.
func (c *Config) View(client *vela.Client) error {
	// send API call to capture a repository
	repo, _, err := client.Repo.Get(c.Org, c.Name)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the repository in dump format
		err := output.Dump(repo)
		if err != nil {
			return err
		}
	case output.DriverJSON:
		// output the repository in JSON format
		err := output.JSON(repo)
		if err != nil {
			return err
		}
	case output.DriverSpew:
		// output the repository in spew format
		err := output.Spew(repo)
		if err != nil {
			return err
		}
	case output.DriverYAML:
		// output the repository in YAML format
		err := output.YAML(repo)
		if err != nil {
			return err
		}
	default:
		// output the repository in stdout format
		err := output.Stdout(repo)
		if err != nil {
			return err
		}
	}

	return nil
}
