// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"
)

// Get captures a list of builds based off the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	logrus.Debug("executing get for build configuration")

	// set the pagination options for list of builds
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ListOptions
	opts := &vela.BuildListOptions{
		Branch: c.Branch,
		Event:  c.Event,
		Status: c.Status,
		ListOptions: vela.ListOptions{
			Page:    c.Page,
			PerPage: c.PerPage,
		},
	}

	logrus.Tracef("capturing builds for repo %s/%s", c.Org, c.Repo)

	// send API call to capture a list of builds
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#BuildService.GetAll
	builds, _, err := client.Build.GetAll(c.Org, c.Repo, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the builds in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(builds)
	case output.DriverJSON:
		// output the builds in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(builds)
	case output.DriverSpew:
		// output the builds in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(builds)
	case "wide":
		// output the builds in wide table format
		return wideTable(builds)
	case output.DriverYAML:
		// output the builds in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(builds)
	default:
		// output the builds in table format
		return table(builds)
	}
}
