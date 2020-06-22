// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-vela/compiler/compiler"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	// check if pipeline action is generate
	if c.Action == "generate" {
		if len(c.File) == 0 {
			return fmt.Errorf("no pipeline file provided")
		}
	}

	return nil
}

// ValidateFile verifies a pipeline based off the provided configuration.
func (c *Config) ValidateFile(client compiler.Engine) error {
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

	// parse the object into a pipeline
	pipeline, err := client.Parse(path)
	if err != nil {
		return err
	}

	// validate the pipeline
	err = client.Validate(pipeline)
	if err != nil {
		return err
	}

	fmt.Printf("%s is valid\n", path)

	return nil
}
