// SPDX-License-Identifier: Apache-2.0

package service

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// View inspects a service based on the provided configuration.
func (c *Config) View(client *vela.Client) error {
	logrus.Debug("executing view for service configuration")

	logrus.Tracef("inspecting service %s/%s/%d/%d", c.Org, c.Repo, c.Build, c.Number)

	// send API call to capture a service
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#SvcService.Get
	service, _, err := client.Svc.Get(c.Org, c.Repo, c.Build, c.Number)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the service in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(service)
	case output.DriverJSON:
		// output the service in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(service)
	case output.DriverSpew:
		// output the service in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(service)
	case output.DriverYAML:
		// output the service in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(service, c.Color)
	default:
		// output the service in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(service)
	}
}
