// SPDX-License-Identifier: Apache-2.0

package log

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of build logs based on the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	logrus.Debug("executing get for log configuration")

	logrus.Tracef("capturing logs for build %s/%s/%d", c.Org, c.Repo, c.Build)

	// create list options for logs call
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	// send API call to capture a list of build logs
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#BuildService.GetLogs
	logs, _, err := client.Build.GetLogs(c.Org, c.Repo, c.Build, opts)
	if err != nil {
		return err
	}

	// create variable for storing all build logs
	data := []byte{}

	// iterate through all build logs
	for _, log := range *logs {
		// add the logs for the step from the build
		data = append(data, log.GetData()...)

		// add a new line to separate the logs
		data = append(data, []byte("\n")...)
	}

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
		return output.YAML(logs, c.Color)
	default:
		// output the logs in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(string(data))
	}
}
