// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/types/library"
	"github.com/go-vela/types/yaml"

	"github.com/go-vela/server/compiler"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating pipeline configuration")

	// handle the action based off the provided configuration
	switch c.Action {
	case internal.ActionCompile:
		fallthrough
	case internal.ActionExpand:
		fallthrough
	case internal.ActionView:
		// check if pipeline org is set
		if len(c.Org) == 0 {
			return fmt.Errorf("no pipeline org provided")
		}

		// check if pipeline repo is set
		if len(c.Repo) == 0 {
			return fmt.Errorf("no pipeline name provided")
		}
	case internal.ActionGenerate:
		fallthrough
	case internal.ActionValidate:
		if len(c.Org) == 0 || len(c.Repo) == 0 {
			// check if pipeline file is set
			if len(c.File) == 0 {
				return fmt.Errorf("no pipeline file provided")
			}
		}

		for _, file := range c.TemplateFiles {
			parts := strings.Split(file, ":")

			// nolint:gomnd,lll // ignore magic number and line length we are explicitly checking for it to parsed into two parts only
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

	logrus.Tracef("parsing pipeline %s", path)

	// set pipelineType within client
	client.WithRepo(&library.Repo{PipelineType: &c.PipelineType})

	// parse the object into a pipeline
	p, err := client.Parse(path)
	if err != nil {
		return err
	}

	templates := mapFromTemplates(p.Templates)

	logrus.Tracef("validating pipeline %s", path)

	if c.Template {
		logrus.Tracef("expand pipeline %s", path)

		// count all the templates
		nTemplates := len(templates)
		nTemplateFiles := len(c.TemplateFiles)

		// ensure a 'templates' block exists in the pipeline
		if nTemplates == 0 {
			return fmt.Errorf("templates block not properly configured in pipeline")
		}

		// whole client put into local mode when we define any template files, override templates locally
		if nTemplateFiles > 0 && nTemplateFiles < nTemplates {
			// nolint:lll // helpful error messages breaks line length
			return fmt.Errorf("found %d template references in your pipeline, but only %d template(s) given to override", nTemplates, nTemplateFiles)
		}

		for _, file := range c.TemplateFiles {
			// local templates override format is <name>:<source>, example: example:/path/to/template.yml
			parts := strings.Split(file, ":")

			// make sure the template was configured
			if _, ok := templates[parts[0]]; !ok {
				return fmt.Errorf("template with name %q is not configured", parts[0])
			}

			// override the source for the given template
			templates[parts[0]].Source = parts[1]
		}

		if len(p.Stages) > 0 {
			// inject the templates into the stages
			p.Stages, p.Secrets, p.Services, _, err = client.ExpandStages(p, templates)
			if err != nil {
				return err
			}
		}

		// inject the templates into the steps
		p.Steps, p.Secrets, p.Services, _, err = client.ExpandSteps(p, templates)
		if err != nil {
			return err
		}
	}

	// validate the pipeline
	if err = client.Validate(p); err != nil {
		return err
	}

	// output the message in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	if err = output.Stdout(fmt.Sprintf("%s is valid", path)); err != nil {
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
		Ref:      c.Ref,
		Template: c.Template,
	}

	// send API call to validate a pipeline
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#PipelineService.Validate
	pipeline, _, err := client.Pipeline.Validate(c.Org, c.Repo, opts)
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

// helper function that creates a map of templates from a yaml configuration.
func mapFromTemplates(templates []*yaml.Template) map[string]*yaml.Template {
	m := make(map[string]*yaml.Template)

	for _, tmpl := range templates {
		m[tmpl.Name] = tmpl
	}

	return m
}
