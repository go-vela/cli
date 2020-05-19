// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"encoding/json"
	"fmt"

	"github.com/go-vela/sdk-go/vela"
)

// Build represents the configuration necessary
// to perform build related quests with Vela.
type Build struct {
	Action  string
	Org     string
	Repo    string
	Number  int
	Page    int
	PerPage int
	Output  string
}

// Get captures a list of builds based off the provided configuration.
func (b *Build) Get(client *vela.Client) error {
	// set the pagination options for list of builds
	opts := &vela.ListOptions{
		Page:    b.Page,
		PerPage: b.PerPage,
	}

	// send API call to capture a list of builds
	builds, _, err := client.Build.GetAll(b.Org, b.Repo, opts)
	if err != nil {
		return err
	}

	switch b.Output {
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

// Restart restarts a build based off the provided configuration.
func (b *Build) Restart(client *vela.Client) error {
	// send API call to restart a build
	build, _, err := client.Build.Restart(b.Org, b.Repo, b.Number)
	if err != nil {
		return err
	}

	switch b.Output {
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

// Validate verifies the Build is properly configured.
func (b *Build) Validate() error {
	// check if build org is set
	if len(b.Org) == 0 {
		return fmt.Errorf("no build org provided")
	}

	// check if build repo is set
	if len(b.Repo) == 0 {
		return fmt.Errorf("no build repo provided")
	}

	// check if build action is restart or view
	if b.Action == restartAction || b.Action == viewAction {
		// check if build number is set
		if b.Number <= 0 {
			return fmt.Errorf("no build number provided")
		}
	}

	return nil
}

// View inspects a build based off the provided configuration.
func (b *Build) View(client *vela.Client) error {
	// send API call to capture a build
	build, _, err := client.Build.Get(b.Org, b.Repo, b.Number)
	if err != nil {
		return err
	}

	switch b.Output {
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
