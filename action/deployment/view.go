// SPDX-License-Identifier: Apache-2.0

package deployment

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// View inspects a deployment based off the provided configuration.
func (c *Config) View(client *vela.Client) error {
	logrus.Debug("executing view for deployment configuration")

	logrus.Tracef("inspecting deployment %s/%s/%d", c.Org, c.Repo, c.Number)

	// send API call to capture a deployment
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#DeploymentService.Get
	deployment, _, err := client.Deployment.Get(c.Org, c.Repo, c.Number)
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
		return output.JSON(deployment, c.Color)
	case output.DriverSpew:
		// output the deployment in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(deployment)
	case output.DriverYAML:
		// output the deployment in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(deployment, c.Color)
	default:
		// output the deployment in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(deployment)
	}
}
