// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"
)

// Expand expands a pipeline based off the provided configuration.
func (c *Config) Expand(client *vela.Client) error {
	logrus.Debug("executing expand for pipeline configuration")

	logrus.Tracef("expanding pipeline %s/%s@%s", c.Org, c.Repo, c.Ref)

	// send API call to expand a pipeline
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#PipelineService.Expand
	pipeline, _, err := client.Pipeline.Expand(c.Org, c.Repo)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the pipeline in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(pipeline)
	case output.DriverJSON:
		// output the pipeline in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(pipeline)
	case output.DriverSpew:
		// output the pipeline in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(pipeline)
	case output.DriverYAML:
		// output the pipeline in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(pipeline)
	default:
		// output the pipeline in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(pipeline)
	}
}
