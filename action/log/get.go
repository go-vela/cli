// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"
)

// Get captures a list of build logs based on the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	logrus.Debug("executing get for log configuration")

	// send API call to capture a list of build logs
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#BuildService.GetLogs
	logs, _, err := client.Build.GetLogs(c.Org, c.Repo, c.Build)
	if err != nil {
		return err
	}

	logrus.Tracef("capturing logs for build %s/%s/%d", c.Org, c.Repo, c.Build)

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the logs in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(logs)
	case output.DriverJSON:
		// output the logs in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(logs)
	case output.DriverSpew:
		// output the logs in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(logs)
	case output.DriverYAML:
		// output the logs in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(logs)
	default:
		// output the logs in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(logs)
	}
}
