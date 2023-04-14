// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package worker

import (
	"fmt"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"
)

// Add creates a worker based off the provided configuration.
func (c *Config) Add(client *vela.Client) error {
	logrus.Debug("executing add for worker configuration")

	// send API call to get a registration token for the given worker
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela#AdminWorkerService.RegisterToken
	registerToken, _, err := client.Admin.Worker.RegisterToken(c.Hostname)
	if err != nil || registerToken == nil {
		return fmt.Errorf("unable to retrieve registration token: %w", err)
	}

	logrus.Tracef("got registration token, adding worker %q", c.Hostname)

	// send API call to register a worker
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela#AdminWorkerService.Register
	msg, _, err := client.Admin.Worker.Register(c.Address, registerToken.GetToken())
	if err != nil {
		return fmt.Errorf("unable to register worker at %q: %w", c.Address, err)
	}

	logrus.Tracef("worker %q registered", c.Hostname)

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the worker in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(msg)
	case output.DriverJSON:
		// output the worker in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(msg)
	case output.DriverSpew:
		// output the worker in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(msg)
	case output.DriverYAML:
		// output the worker in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(msg)
	default:
		// output the worker in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(*msg)
	}
}
