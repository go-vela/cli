// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"fmt"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"
)

// Get captures a list of workers based off the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	logrus.Debug("executing get for worker configuration")

	logrus.Tracef("capturing workers")

	filters := &vela.WorkerListOptions{
		CheckedInAfter: c.CheckedInAfter,
	}

	if c.Active != nil {
		filters.Active = fmt.Sprintf("%b", c.Active)
	}

	if c.CheckedInBefore > 0 {
		filters.CheckedInBefore = c.CheckedInBefore
	}

	// send API call to capture a list of workers
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#WorkerService.GetAll
	workers, _, err := client.Worker.GetAll(filters)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the workers in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(workers)
	case output.DriverJSON:
		// output the workers in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(workers)
	case output.DriverSpew:
		// output the workers in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(workers)
	case "wide":
		// output the workers in wide table format
		return wideTable(workers)
	case output.DriverYAML:
		// output the workers in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(workers)
	default:
		// output the workers in table format
		return table(workers)
	}
}
