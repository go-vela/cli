// SPDX-License-Identifier: Apache-2.0

package service

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of services based on the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	logrus.Debug("executing get for service configuration")

	// set the pagination options for list of services
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ListOptions
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	logrus.Tracef("capturing services for build %s/%s/%d", c.Org, c.Repo, c.Build)

	// send API call to capture a list of services
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#SvcService.GetAll
	services, _, err := client.Svc.GetAll(c.Org, c.Repo, c.Build, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the services in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(services)
	case output.DriverJSON:
		// output the services in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(services, c.Color)
	case output.DriverSpew:
		// output the services in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(services)
	case "wide":
		// output the services in wide table format
		return wideTable(services)
	case output.DriverYAML:
		// output the services in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(services, c.Color)
	default:
		// output the services in table format
		return table(services)
	}
}
