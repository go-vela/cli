// SPDX-License-Identifier: Apache-2.0

package schedule

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Remove deletes a schedule based off the provided configuration.
func (c *Config) Remove(client *vela.Client) error {
	logrus.Debug("executing remove for schedule configuration")

	logrus.Tracef("removing schedule %s/%s/%s", c.Org, c.Repo, c.Name)

	// send API call to remove a schedule
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ScheduleService.Remove
	msg, _, err := client.Schedule.Remove(c.Org, c.Repo, c.Name)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the message in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(msg)
	case output.DriverJSON:
		// output the message in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(msg, c.Color)
	case output.DriverSpew:
		// output the message in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(msg)
	case output.DriverYAML:
		// output the message in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(msg, c.Color)
	default:
		// output the message in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(*msg)
	}
}
