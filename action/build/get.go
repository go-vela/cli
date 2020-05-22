// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"encoding/json"
	"fmt"

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

	switch c.Output {
	case "json":
		// TODO: create output package
		//
		// err := output.JSON(builds)
		// if err != nil {
		// 	return err
		// }

		fallthrough
	case "wide":
		// TODO: create output package
		//
		// err := output.Wide(builds)
		// if err != nil {
		// 	return err
		// }

		fallthrough
	case "yaml":
		// TODO: create output package
		//
		// err := output.YAML(builds)
		// if err != nil {
		// 	return err
		// }

		fallthrough
	default:
		// TODO: create output package
		//
		// err := output.Default(builds)
		// if err != nil {
		// 	return err
		// }

		output, err := json.MarshalIndent(builds, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
	}

	return nil
}
