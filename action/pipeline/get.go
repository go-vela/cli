// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of pipelines based on the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	logrus.Debug("executing get for pipeline configuration")

	// set the pagination options for list of pipelines
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ListOptions
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	logrus.Tracef("capturing pipelines for repo %s/%s", c.Org, c.Repo)

	// send API call to capture a list of pipelines
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#PipelineService.GetAll
	pipelines, _, err := client.Pipeline.GetAll(c.Org, c.Repo, opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the pipelines in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(pipelines)
	case output.DriverJSON:
		// output the pipelines in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(pipelines)
	case output.DriverSpew:
		// output the pipelines in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(pipelines)
	case "wide":
		// output the pipelines in wide table format
		return wideTable(pipelines)
	case output.DriverYAML:
		// output the pipelines in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(pipelines)
	default:
		// output the pipelines in table format
		return table(pipelines)
	}
}
