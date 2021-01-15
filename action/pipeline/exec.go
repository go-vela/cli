// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-vela/compiler/compiler"
	"github.com/go-vela/pkg-executor/executor"
	"github.com/go-vela/pkg-runtime/runtime"
	"github.com/go-vela/types/constants"

	"github.com/sirupsen/logrus"
)

// Exec executes a pipeline based off the provided configuration.
//
// nolint: funlen // ignore function length due to comments
func (c *Config) Exec(client compiler.Engine) error {
	logrus.Debug("executing exec for pipeline configuration")

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

	logrus.Tracef("compiling pipeline %s", path)

	// compile into a pipeline
	_pipeline, err := client.
		WithLocal(true).
		Compile(path)
	if err != nil {
		return err
	}

	// check if the local configuration is enabled
	if c.Local {
		// TODO: consider making this a constant
		//
		// create current directory path for local mount
		mount := fmt.Sprintf("%s:%s", base, "/vela/src")

		// add the current directory path to volume mounts
		c.Volumes = append(c.Volumes, mount)
	}

	logrus.Tracef("creating runtime engine %s", constants.DriverDocker)

	// setup the runtime
	//
	// https://pkg.go.dev/github.com/go-vela/pkg-runtime/runtime?tab=doc#New
	_runtime, err := runtime.New(&runtime.Setup{
		Driver:  constants.DriverDocker,
		Volumes: c.Volumes,
	})
	if err != nil {
		return err
	}

	logrus.Tracef("creating executor engine %s", constants.DriverLocal)

	// setup the executor
	//
	// https://godoc.org/github.com/go-vela/pkg-executor/executor#New
	_executor, err := executor.New(&executor.Setup{
		// TODO: switch to using constant from go-vela/types
		Driver:   constants.DriverLocal,
		Runtime:  _runtime,
		Pipeline: _pipeline.Sanitize(constants.DriverDocker),
	})
	if err != nil {
		return err
	}

	// create a background context
	ctx := context.Background()

	defer func() {
		// destroy the build with the executor
		err = _executor.DestroyBuild(context.Background())
		if err != nil {
			logrus.Errorf("unable to destroy build: %v", err)
		}
	}()

	// create the build with the executor
	err = _executor.CreateBuild(ctx)
	if err != nil {
		return fmt.Errorf("unable to create build: %v", err)
	}

	// plan the build with the executor
	err = _executor.PlanBuild(ctx)
	if err != nil {
		return fmt.Errorf("unable to plan build: %v", err)
	}

	// assemble the build with the executor
	err = _executor.AssembleBuild(ctx)
	if err != nil {
		return fmt.Errorf("unable to assemble build: %v", err)
	}

	// execute the build with the executor
	err = _executor.ExecBuild(ctx)
	if err != nil {
		return fmt.Errorf("unable to execute build: %v", err)
	}

	return nil
}
