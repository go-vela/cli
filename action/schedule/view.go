// SPDX-License-Identifier: Apache-2.0

package schedule

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// View inspects a schedule based off the provided configuration.
func (c *Config) View(client *vela.Client) error {
	logrus.Debug("executing view for schedule configuration")

	logrus.Tracef("inspecting schedule %s/%s/%s", c.Org, c.Repo, c.Name)

	// send API call to capture a schedule
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ScheduleService.Get
	schedule, _, err := client.Schedule.Get(c.Org, c.Repo, c.Name)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the schedule in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(schedule)
	case output.DriverJSON:
		// output the schedule in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(schedule)
	case output.DriverSpew:
		// output the schedule in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(schedule)
	case output.DriverYAML:
		// output the schedule in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(schedule)
	default:
		// output the schedule in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(schedule)
	}
}
