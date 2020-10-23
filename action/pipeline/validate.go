// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/compiler/compiler"

	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating pipeline configuration")

	// check if pipeline action is not view
	if c.Action != "view" {
		// check if pipeline file is set
		if len(c.File) == 0 {
			return fmt.Errorf("no pipeline file provided")
		}

		return nil
	}

	// check if pipeline org is set
	if len(c.Org) == 0 {
		return fmt.Errorf("no pipeline org provided")
	}

	// check if pipeline name is set
	if len(c.Repo) == 0 {
		return fmt.Errorf("no pipeline name provided")
	}

	return nil
}

// ValidateFile verifies a pipeline based off the provided configuration.
func (c *Config) ValidateFile(client compiler.Engine) error {
	logrus.Debug("executing validate for pipeline configuration")

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

	logrus.Tracef("parsing pipeline %s", path)

	// parse the object into a pipeline
	pipeline, err := client.Parse(path)
	if err != nil {
		return err
	}

	logrus.Tracef("validating pipeline %s", path)

	// validate the pipeline
	err = client.Validate(pipeline)
	if err != nil {
		return err
	}

	// output the message in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(fmt.Sprintf("%s is valid", path))
}
