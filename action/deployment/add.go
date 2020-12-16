// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"fmt"
	"strings"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/library"
	"github.com/go-vela/types/raw"

	"github.com/sirupsen/logrus"
)

// parseKeyValue converts the slice of key=value into a map
func parseKeyValue(input []string) (raw.StringSliceMap, error) {
	payload := raw.StringSliceMap{}
	for _, i := range input {
		parts := strings.SplitN(i, "=", 2)
		if len(parts) != 2 { //nolint
			return nil, fmt.Errorf("%s is not in key=value format", i)
		}
		payload[parts[0]] = parts[1]
	}
	return payload, nil
}

// Add creates a deployment based off the provided configuration.
func (c *Config) Add(client *vela.Client) error {
	logrus.Debug("executing add for deployment configuration")

	// create the deployment object
	//
	// https://pkg.go.dev/github.com/go-vela/types/library?tab=doc#Deployment
	d := &library.Deployment{
		Ref:         &c.Ref,
		Task:        &c.Task,
		Target:      &c.Target,
		Description: &c.Description,
	}

	// check if the user provided any parameters
	if len(c.Parameters) != 0 {
		// convert the parameters into map[string]string format
		payload, err := parseKeyValue(c.Parameters)
		if err != nil {
			return err
		}
		d.SetPayload(payload)
	}

	logrus.Tracef("adding deployment for repo %s/%s", c.Org, c.Repo)

	// send API call to add a deployment
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#DeploymentService.Add
	deployment, _, err := client.Deployment.Add(c.Org, c.Repo, d)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the deployment in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(deployment)
	case output.DriverJSON:
		// output the deployment in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(deployment)
	case output.DriverSpew:
		// output the deployment in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(deployment)
	case output.DriverYAML:
		// output the deployment in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(deployment)
	default:
		// output the deployment in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(deployment)
	}
}
