// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package worker

import (
	"fmt"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/library"

	"github.com/sirupsen/logrus"
)

// Add creates a worker based off the provided configuration.
func (c *Config) Add(client *vela.Client) error {
	logrus.Debug("executing add for worker configuration")

	// create the worker object
	//
	// https://pkg.go.dev/github.com/go-vela/types/library?tab=doc#Worker
	w := &library.Worker{
		Hostname: vela.String(c.Hostname),
		Address:  vela.String(c.Address),
	}

	registerToken, _, err := client.Admin.Worker.RegisterToken(c.Hostname)
	if err != nil || registerToken == nil {
		return fmt.Errorf("unable to retrieve registration token: %w", err)
	}

	// set the registration token as the authentication
	// for the next request to add a worker
	client.Authentication.SetTokenAuth(*registerToken.Token)

	logrus.Tracef("adding worker %s", c.Hostname)

	// send API call to add a worker
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#WorkerService.Add
	worker, _, err := client.Worker.Add(w)
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
		return output.JSON(worker)
	case output.DriverSpew:
		// output the worker in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(worker)
	case output.DriverYAML:
		// output the worker in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(worker)
	default:
		// output the worker in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(worker)
	}
}
