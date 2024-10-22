// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
	api "github.com/go-vela/server/api/types"
)

// Update modifies a worker based off the provided configuration.
func (c *Config) Update(client *vela.Client) error {
	logrus.Debug("executing update for worker configuration")

	// create the worker object
	w := &api.Worker{
		Hostname:   vela.String(c.Hostname),
		Address:    vela.String(c.Address),
		Active:     c.Active,
		Routes:     vela.Strings(c.Routes),
		BuildLimit: vela.Int64(c.BuildLimit),
	}

	logrus.Tracef("updating worker %s", c.Hostname)

	// send API call to modify a worker
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#WorkerService.Update
	worker, _, err := client.Worker.Update(c.Hostname, w)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the worker in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(worker)
	case output.DriverJSON:
		// output the worker in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(worker, c.Color)
	case output.DriverSpew:
		// output the worker in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(worker)
	case output.DriverYAML:
		// output the worker in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(worker, c.Color)
	default:
		// output the worker in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(worker)
	}
}
