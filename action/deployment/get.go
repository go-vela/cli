// SPDX-License-Identifier: Apache-2.0

package deployment

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of deployments based off the provided configuration.
func (c *Config) Get(ctx context.Context, client *vela.Client) error {
	logrus.Debug("executing get for deployment configuration")

	// set the pagination options for list of deployments
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ListOptions
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	logrus.Tracef("capturing deployments for repo %s/%s", c.Org, c.Repo)

	// send API call to capture a list of deployments
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#DeploymentService.GetAll
	deployments, _, err := client.Deployment.GetAll(ctx, c.Org, c.Repo, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the deployments in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(deployments)
	case output.DriverJSON:
		// output the deployments in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(deployments, c.Color)
	case output.DriverSpew:
		// output the deployments in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(deployments)
	case "wide":
		// output the deployments in wide table format
		return wideTable(deployments)
	case output.DriverYAML:
		// output the deployments in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(deployments, c.Color)
	default:
		// output the deployments in table format
		return table(deployments)
	}
}
