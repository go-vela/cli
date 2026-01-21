// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// View inspects a worker based off the provided configuration.
// Based on the configuration, it will either return details of
// a worker or view the registration token for the given worker.
func (c *Config) View(ctx context.Context, client *vela.Client) error {
	logrus.Debug("executing view for worker configuration")

	logrus.Tracef("inspecting worker with hostname %s", c.Hostname)

	var (
		response any
		err      error
	)

	// handle RegistrationToken flag
	if c.RegistrationToken {
		response, _, err = client.Admin.Worker.RegisterToken(ctx, c.Hostname)
		if err != nil {
			return fmt.Errorf("unable to retrieve registration token for worker %q: %w", c.Hostname, err)
		}
	} else {
		response, _, err = client.Worker.Get(ctx, c.Hostname)
		if err != nil {
			return fmt.Errorf("unable to retrieve information for worker %q: %w", c.Hostname, err)
		}
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the worker in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(response)
	case output.DriverJSON:
		// output the worker in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(response, c.Color)
	case output.DriverSpew:
		// output the worker in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(response)
	case output.DriverYAML:
		// output the worker in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(response, c.Color)
	default:
		// output the worker in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(response)
	}
}
