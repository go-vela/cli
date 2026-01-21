// SPDX-License-Identifier: Apache-2.0

package schedule

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of schedules based off the provided configuration.
func (c *Config) Get(ctx context.Context, client *vela.Client) error {
	logrus.Debug("executing get for schedule configuration")

	// set the pagination options for list of schedules
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ListOptions
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	logrus.Tracef("capturing schedules for repo %s/%s", c.Org, c.Repo)

	// send API call to capture a list of schedules
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ScheduleService.GetAll
	schedules, _, err := client.Schedule.GetAll(ctx, c.Org, c.Repo, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the schedules in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(schedules)
	case output.DriverJSON:
		// output the schedules in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(schedules, c.Color)
	case output.DriverSpew:
		// output the schedules in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(schedules)
	case "wide":
		// output the schedules in wide table format
		return wideTable(schedules)
	case output.DriverYAML:
		// output the schedules in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(schedules, c.Color)
	default:
		// output the schedules in table format
		return table(schedules)
	}
}
