// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"encoding/json"
	"fmt"

	"github.com/go-vela/sdk-go/vela"
)

// Restart restarts a build based off the provided configuration.
func (c *Config) Restart(client *vela.Client) error {
	// send API call to restart a build
	build, _, err := client.Build.Restart(c.Org, c.Repo, c.Number)
	if err != nil {
		return err
	}

	switch c.Output {
	case "json":
		// TODO: create output package
		//
		// err := output.JSON(build)
		// if err != nil {
		// 	return err
		// }

		fallthrough
	case "yaml":
		// TODO: create output package
		//
		// err := output.YAML(build)
		// if err != nil {
		// 	return err
		// }

		fallthrough
	default:
		// TODO: create output package
		//
		// err := output.Default(build)
		// if err != nil {
		// 	return err
		// }

		output, err := json.MarshalIndent(build, "", "    ")
		if err != nil {
			return err
		}

		fmt.Println(string(output))
	}

	return nil
}
