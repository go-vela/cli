// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/library"

	"github.com/go-vela/server/compiler"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating pipeline configuration")

	// handle the action based off the provided configuration
	switch c.Action {
	case "get":
		// check if pipeline org is set
		if len(c.Org) == 0 {
			return fmt.Errorf("no pipeline org provided")
		}

		// check if pipeline repo is set
		if len(c.Repo) == 0 {
			return fmt.Errorf("no pipeline name provided")
		}
	case "compile":
		fallthrough
	case "expand":
		fallthrough
	case "view":
		// check if pipeline org is set
		if len(c.Org) == 0 {
			return fmt.Errorf("no pipeline org provided")
		}

		// check if pipeline repo is set
		if len(c.Repo) == 0 {
			return fmt.Errorf("no pipeline name provided")
		}

		// check if pipeline ref is set
		if len(c.Ref) == 0 {
			return fmt.Errorf("no pipeline ref provided")
		}
	case "generate":
		fallthrough
	case "validate":
		if len(c.Org) == 0 || len(c.Repo) == 0 {
			// check if pipeline file is set
			if len(c.File) == 0 {
				return fmt.Errorf("no pipeline file provided")
			}
		}

		for _, file := range c.TemplateFiles {
			parts := strings.Split(file, ":")

			if len(parts) != 2 {
				return fmt.Errorf("invalid format for template file: %s (valid format: <name>:<source>)", file)
			}
		}
	}

	return nil
}

// ValidateLocal verifies a local pipeline based off the provided configuration.
func (c *Config) ValidateLocal(client compiler.Engine) error {
	logrus.Debug("executing validate for local pipeline configuration")

	// send Filesystem call to capture base directory path
	base, err := os.Getwd()
	if err != nil {
		return err
	}

	// create full path for pipeline file
	path := filepath.Join(base, c.File)

	// check if custom path was provided for pipeline file
	if len(c.Path) > 0 {
		// create custom full path for pipeline file
		path = filepath.Join(c.Path, c.File)
	}

	// check if full path to pipeline file exists
	path, err = validateFile(path)
	if err != nil {
		return err
	}

	// set pipelineType within client
	client.WithRepo(&library.Repo{PipelineType: &c.PipelineType})

	logrus.Tracef("compiling pipeline %s", path)

	// compile the object into a pipeline
	p, _, err := client.CompileLite(path, c.Template, false)
	if err != nil {
		return err
	}

	// output the message in stderr format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stderr
	err = output.Stderr(fmt.Sprintf("%s is valid", path))
	if err != nil {
		return err
	}

	// output the validated pipeline in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.YAML(p)
}

// ValidateRemote validates a remote pipeline based off the provided configuration.
func (c *Config) ValidateRemote(client *vela.Client) error {
	logrus.Debug("executing validate for remote pipeline configuration")

	logrus.Tracef("validating pipeline %s/%s@%s", c.Org, c.Repo, c.Ref)

	// set the pipeline options for the call
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#PipelineOptions
	opts := &vela.PipelineOptions{
		Output:   c.Output,
		Template: c.Template,
	}

	// send API call to validate a pipeline
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#PipelineService.Validate
	pipeline, _, err := client.Pipeline.Validate(c.Org, c.Repo, c.Ref, opts)
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

// validateFile validates the configuration file exists.
func validateFile(path string) (string, error) {
	// check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// attempt to validate if .vela.yaml exists if .vela.yml does not
		if filepath.Base(path) == ".vela.yml" {
			// override path if .vela.yaml exists
			if _, err := os.Stat(filepath.Join(filepath.Dir(path), ".vela.yaml")); err == nil {
				return filepath.Join(filepath.Dir(path), ".vela.yaml"), nil
			}
		}

		return path, fmt.Errorf("configuration file of %s does not exist", path)
	}

	return path, nil
}
