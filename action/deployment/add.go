// SPDX-License-Identifier: Apache-2.0

package deployment

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/library"

	"github.com/sirupsen/logrus"
)

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
		Payload:     &c.Parameters,
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
