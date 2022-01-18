// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// nolint: dupl // ignore similar code among actions
package repo

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"
)

// Repair recreates a repository webhook based off the provided configuration.
func (c *Config) Repair(client *vela.Client) error {
	logrus.Debug("executing repair for repo configuration")

	logrus.Tracef("repairing repo %s/%s", c.Org, c.Name)

	// send API call to repair a repository
	msg, _, err := client.Repo.Repair(c.Org, c.Name)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the message in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(msg)
	case output.DriverJSON:
		// output the message in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(msg)
	case output.DriverSpew:
		// output the message in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(msg)
	case output.DriverYAML:
		// output the message in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(msg)
	default:
		// output the message in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(*msg)
	}
}
