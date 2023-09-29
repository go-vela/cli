// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/go-vela/cli/version"
	"github.com/go-vela/server/compiler"
	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
	"github.com/go-vela/worker/executor"
	"github.com/go-vela/worker/runtime"

	"github.com/sirupsen/logrus"
)

// Exec executes a pipeline based off the provided configuration.
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

	path, err = validateFile(path)
	if err != nil {
		return err
	}

	// check if full path to pipeline file exists
	_, err = os.Stat(path)
	if err != nil {
		return fmt.Errorf("unable to find pipeline %s: %w", path, err)
	}

	// create build object for use in pipeline
	b := new(library.Build)
	b.SetBranch(c.Branch)
	b.SetDeploy(c.Target)
	b.SetEvent(c.Event)

	if c.Tag == "" && c.Event == constants.EventPull {
		b.SetRef("refs/pull/1")
	} else {
		b.SetRef(c.Tag)
	}

	// create repo object for use in pipeline
	r := new(library.Repo)
	r.SetOrg(c.Org)
	r.SetName(c.Repo)
	r.SetFullName(fmt.Sprintf("%s/%s", c.Org, c.Repo))
	r.SetPipelineType(c.PipelineType)

	logrus.Tracef("compiling pipeline %s", path)

	// compile into a pipeline
	_pipeline, _, err := client.
		Duplicate().
		WithBuild(b).
		WithComment(c.Comment).
		WithLocal(true).
		WithRepo(r).
		Compile(path)
	if err != nil {
		return err
	}

	// check if the local configuration is enabled
	if c.Local {
		// create current directory path for local mount
		mount := fmt.Sprintf("%s:%s:rw", base, constants.WorkspaceDefault)

		// add the current directory path to volume mounts
		c.Volumes = append(c.Volumes, mount)
	}

	logrus.Tracef("creating runtime engine %s", constants.DriverDocker)

	// setup the runtime
	//
	// https://pkg.go.dev/github.com/go-vela/worker/runtime?tab=doc#New
	_runtime, err := runtime.New(&runtime.Setup{
		Driver:      constants.DriverDocker,
		HostVolumes: c.Volumes,
	})
	if err != nil {
		return err
	}

	logrus.Tracef("creating executor engine %s", constants.DriverLocal)

	// setup the executor
	//
	// https://godoc.org/github.com/go-vela/worker/executor#New
	_executor, err := executor.New(&executor.Setup{
		Driver:   constants.DriverLocal,
		Runtime:  _runtime,
		Pipeline: _pipeline.Sanitize(constants.DriverDocker),
		Build:    b,
		Repo:     r,
		Version:  version.New().Semantic(),
	})
	if err != nil {
		return err
	}

	// create a background context
	ctx, done := context.WithCancel(context.Background())
	defer done()

	// handle aborting local build process
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// spawn go routine to wait for syscall signals
	go func() {
		// wait for signal
		<-signalChan

		logrus.Info("pipeline exec canceled! cleaning up - you may see some errors during cleanup")

		// cancel the context passed into build process
		done()
	}()

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
		return fmt.Errorf("unable to create build: %w", err)
	}

	// plan the build with the executor
	err = _executor.PlanBuild(ctx)
	if err != nil {
		return fmt.Errorf("unable to plan build: %w", err)
	}

	// log/event streaming
	go func() {
		logrus.Debug("streaming build logs")
		// start process to handle StreamRequests
		// from Steps and Services
		err = _executor.StreamBuild(ctx)
		if err != nil {
			logrus.Errorf("unable to stream build logs: %v", err)
		}
	}()

	// assemble the build with the executor
	err = _executor.AssembleBuild(ctx)
	if err != nil {
		return fmt.Errorf("unable to assemble build: %w", err)
	}

	// execute the build with the executor
	err = _executor.ExecBuild(ctx)
	if err != nil {
		return fmt.Errorf("unable to execute build: %w", err)
	}

	return nil
}
