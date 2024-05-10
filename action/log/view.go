// SPDX-License-Identifier: Apache-2.0

package log

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// ViewService inspects a service log based on the provided configuration.
//
//nolint:dupl // ignore similar code among actions
func (c *Config) ViewService(client *vela.Client) error {
	logrus.Debug("executing view service for log configuration")

	logrus.Tracef("capturing logs for service %s/%s/%d", c.Org, c.Repo, c.Service)

	// send API call to capture a service log
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#LogService.GetService
	log, _, err := client.Log.GetService(c.Org, c.Repo, c.Build, c.Service)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the service log in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(log)
	case output.DriverJSON:
		// output the service log in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(log)
	case output.DriverSpew:
		// output the service log in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(log)
	case output.DriverYAML:
		// output the service log in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(log, c.Color)
	default:
		// output the service log in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(string(log.GetData()))
	}
}

// ViewStep inspects a service log based on the provided configuration.
//
//nolint:dupl // ignore similar code among actions
func (c *Config) ViewStep(client *vela.Client) error {
	logrus.Debug("executing view step for log configuration")

	logrus.Tracef("capturing logs for step %s/%s/%d", c.Org, c.Repo, c.Step)

	// send API call to capture a step log
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#LogService.GetStep
	log, _, err := client.Log.GetStep(c.Org, c.Repo, c.Build, c.Step)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the step log in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(log)
	case output.DriverJSON:
		// output the step log in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(log)
	case output.DriverSpew:
		// output the step log in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(log)
	case output.DriverYAML:
		// output the step log in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(log, c.Color)
	default:
		// output the step log in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(string(log.GetData()))
	}
}
