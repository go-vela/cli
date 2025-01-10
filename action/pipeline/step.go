// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/server/compiler/types/yaml/yaml"
)

func steps(pipelineType string) *yaml.Build {
	logrus.Debugf("creating %s steps pipeline", pipelineType)

	// create default image for steps pipeline
	image := "alpine:latest"

	// create default commands for steps pipeline
	commands := []string{"echo hello"}

	// handle the pipeline based off the provided type
	switch pipelineType {
	case "go", "golang":
		// set the image for a go steps pipeline
		image = "golang:latest"
		// set the commands for a go steps pipeline
		commands = []string{"go version"}
	case "java":
		// set the image for a java steps pipeline
		image = "openjdk:latest"
		// set the commands for a java steps pipeline
		commands = []string{"java --version"}
	case "node", "node.js":
		// set the image for a node steps pipeline
		image = "node:latest"
		// set the commands for a node steps pipeline
		commands = []string{"node --version"}
	}

	// return a steps pipeline based off the type
	//
	// https://pkg.go.dev/github.com/go-vela/server/compiler/types/yaml?tab=doc#Build
	return &yaml.Build{
		Version: "1",
		Steps: yaml.StepSlice{
			{
				Commands: commands,
				Image:    image,
				Name:     "version",
				Pull:     "always",
			},
		},
	}
}
