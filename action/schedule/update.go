// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code among actions
package schedule

import (
	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/library"
	"github.com/sirupsen/logrus"
)

// Update modifies a schedule based off the provided configuration.
func (c *Config) Update(client *vela.Client) error {
	logrus.Debug("executing update for schedule configuration")

	// create the schedule object
	s := &library.Schedule{
		Active: vela.Bool(c.Active),
		Name:   vela.String(c.Name),
		Entry:  vela.String(c.Entry),
		Branch: vela.String(c.Branch),
	}

	logrus.Tracef("updating schedule %s/%s/%s", c.Org, c.Repo, c.Name)

	// send API call to modify a schedule
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ScheduleService.Update
	schedule, _, err := client.Schedule.Update(c.Org, c.Repo, s)
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
