// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"github.com/go-vela/types/yaml"
)

func stages(pipelineType string) *yaml.Build {
	// create default image for stages pipeline
	image := "alpine:latest"

	// create default commands for stages pipeline
	commands := []string{"echo hello"}

	// handle the pipeline based off the provided type
	switch pipelineType {
	case "go", "golang":
		// set the image for a go stages pipeline
		image = "golang:latest"
		// set the commands for a go stages pipeline
		commands = []string{"go version"}
	case "java":
		// set the image for a java stages pipeline
		image = "openjdk:latest"
		// set the commands for a java stages pipeline
		commands = []string{"java --version"}
	case "node", "node.js":
		// set the image for a node stages pipeline
		image = "node:latest"
		// set the commands for a node stages pipeline
		commands = []string{"node --version"}
	}

	// return a stages pipeline based off the type
	return &yaml.Build{
		Version: "1",
		Stages: yaml.StageSlice{
			{
				Name: "version",
				Steps: yaml.StepSlice{
					{
						Commands: commands,
						Image:    image,
						Name:     "version",
					},
				},
			},
		},
	}
}
