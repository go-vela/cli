// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"slices"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/version"
	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/compiler"
	"github.com/go-vela/server/compiler/types/pipeline"
	"github.com/go-vela/server/constants"
	"github.com/go-vela/worker/executor"
	"github.com/go-vela/worker/runtime"
)

// Exec executes a pipeline based off the provided configuration.
//
//nolint:funlen // ignore function length
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
	b := new(api.Build)
	b.SetBranch(c.Branch)
	b.SetDeploy(c.Target)

	fullEvent := strings.Split(c.Event, ":")
	if len(fullEvent) == 2 {
		b.SetEvent(fullEvent[0])
		b.SetEventAction(fullEvent[1])
	} else {
		b.SetEvent(c.Event)

		switch c.Event {
		case constants.EventPull, constants.EventPullAlternate:
			logrus.Debug("setting pull_request event action as `opened`")
			b.SetEvent(constants.EventPull)
			b.SetEventAction(constants.ActionOpened)
		case constants.EventComment:
			logrus.Debug("setting comment event action as `created`")
			b.SetEvent(constants.EventComment)
			b.SetEventAction(constants.ActionCreated)
		case constants.EventDeploy, constants.EventDeployAlternate:
			logrus.Debug("setting deployment event action as `created`")
			b.SetEvent(constants.EventDeploy)
			b.SetEventAction(constants.ActionCreated)
		case constants.EventDelete:
			return fmt.Errorf("event %s must supply an action (branch or tag)", c.Event)
		}
	}

	if c.Tag == "" && b.GetEvent() == constants.EventPull {
		b.SetRef("refs/pull/1")
	} else {
		b.SetRef(c.Tag)
	}

	// create repo object for use in pipeline
	r := new(api.Repo)
	r.SetOrg(c.Org)
	r.SetName(c.Repo)
	r.SetFullName(fmt.Sprintf("%s/%s", c.Org, c.Repo))
	r.SetPipelineType(c.PipelineType)

	b.SetRepo(r)

	logrus.Tracef("compiling pipeline %s", path)

	// compile into a pipeline
	_pipeline, _, err := client.
		Duplicate().
		WithBuild(b).
		WithComment(c.Comment).
		WithFiles(c.FileChangeset).
		WithLocal(true).
		WithRepo(r).
		WithLocalTemplates(c.TemplateFiles).
		Compile(context.Background(), path)
	if err != nil {
		return err
	}

	// create a slice for steps to be removed
	stepsToRemove := c.SkipSteps

	// print and remove steps
	if len(stepsToRemove) > 0 {
		for _, stepName := range stepsToRemove {
			logrus.Info("skipping step: ", stepName)
		}

		if err := skipSteps(_pipeline, stepsToRemove); err != nil {
			return err
		}
	}

	// find all secrets that were not provided
	missingSecrets := collectMissingSecrets(_pipeline)

	// create current directory path for local mount
	mount := fmt.Sprintf("%s:%s:rw", base, constants.WorkspaceDefault)

	// add the current directory path to volume mounts
	c.Volumes = append(c.Volumes, mount)

	logrus.Tracef("creating runtime engine %s", constants.DriverDocker)

	// setup the runtime
	//
	// https://pkg.go.dev/github.com/go-vela/worker/runtime?tab=doc#New
	_runtime, err := runtime.New(&runtime.Setup{
		Driver:           constants.DriverDocker,
		HostVolumes:      c.Volumes,
		PrivilegedImages: c.PrivilegedImages,
	})
	if err != nil {
		return err
	}

	logrus.Tracef("creating executor engine %s", constants.DriverLocal)

	execSetup := &executor.Setup{
		Driver:   constants.DriverLocal,
		Runtime:  _runtime,
		Pipeline: _pipeline.Sanitize(constants.DriverDocker),
		Build:    b,
		Version:  version.New().Semantic(),
	}

	if len(c.OutputsImage) > 0 {
		logrus.Debugf("using image %s for outputs container", c.OutputsImage)

		outputsCtn := &pipeline.Container{
			ID:          fmt.Sprintf("outputs_%s_%s", r.GetOrg(), r.GetName()),
			Detach:      true,
			Image:       c.OutputsImage,
			Environment: make(map[string]string),
			Pull:        constants.PullNotPresent,
		}

		execSetup.OutputCtn = outputsCtn
	}

	_executor, err := executor.New(execSetup)
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
		// print any secrets not set to the user
		reportMissingSecrets(missingSecrets)

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

// reportMissingSecrets informs the user of any secrets not set.
func reportMissingSecrets(s map[string]string) {
	if len(s) > 0 {
		logrus.Warn("the following secrets were not set, use --help to learn how to set them:")

		for step, secret := range s {
			logrus.Warnf("secret %#q not set for step %#q", secret, step)
		}
	}
}

// collectMissingSecrets searches a given pipeline for used secrets
// and returns a map of secrets not set in the current environment.
// The map key is is the step, stage+step, or secret name formatted
// to match the local exec log output.
func collectMissingSecrets(p *pipeline.Build) map[string]string {
	if p == nil {
		return make(map[string]string)
	}

	secrets := map[string]string{}

	for _, stage := range p.Stages {
		for _, step := range stage.Steps {
			for _, secret := range step.Secrets {
				stepName := formatStepIdentifier(stage.Name, step.Name, false)
				secrets[stepName] = secret.Target
			}
		}
	}

	for _, step := range p.Steps {
		for _, secret := range step.Secrets {
			stepName := formatStepIdentifier("", step.Name, false)
			secrets[stepName] = secret.Target
		}
	}

	for _, s := range p.Secrets {
		if !s.Origin.Empty() {
			for _, secret := range s.Origin.Secrets {
				stepName := formatStepIdentifier("", s.Origin.Name, true)
				secrets[stepName] = secret.Target
			}
		}
	}

	for step, secret := range secrets {
		// if the secret was supplied, remove it from the map
		// we only care about unset secrets
		val, exists := os.LookupEnv(secret)
		if exists {
			delete(secrets, step)

			if val == "" {
				logrus.Debugf("secret %#q for step %#q is provided but empty", secret, step)
			}
		}
	}

	return secrets
}

// formatStepIdentifier formats a step name to be consistent with what
// the worker logs to make it easier to associate a missing secret
// with a step.
func formatStepIdentifier(stageName, stepName string, isSecret bool) string {
	const (
		secretPrefix = "[secret: %s]" //nolint:gosec // false positive
		stagePrefix  = "[stage: %s]"
		stepPrefix   = "[step: %s]"
	)

	output := strings.Builder{}

	if stageName != "" {
		output.WriteString(fmt.Sprintf(stagePrefix, stageName))
	}

	if stepName != "" {
		if isSecret {
			output.WriteString(fmt.Sprintf(secretPrefix, stepName))
		} else {
			output.WriteString(fmt.Sprintf(stepPrefix, stepName))
		}
	}

	return output.String()
}

// skipSteps filters out steps to be removed from the pipeline.
func skipSteps(_pipeline *pipeline.Build, stepsToRemove []string) error {
	// filter out steps to be removed
	if len(_pipeline.Stages) > 0 {
		// counter for total steps to run
		totalSteps := 0

		for i, stage := range _pipeline.Stages {
			filteredStageSteps := stage.Steps[:0]

			for _, step := range stage.Steps {
				if !slices.Contains(stepsToRemove, step.Name) {
					filteredStageSteps = append(filteredStageSteps, step)
					totalSteps++
				}
			}

			_pipeline.Stages[i].Steps = filteredStageSteps
		}

		// check if any steps are left to run, excluding "init" step
		if totalSteps <= 1 {
			return fmt.Errorf("no steps left to run after removing skipped steps")
		}
	} else {
		// if not using stages
		filteredSteps := _pipeline.Steps[:0]

		for _, step := range _pipeline.Steps {
			if !slices.Contains(stepsToRemove, step.Name) {
				filteredSteps = append(filteredSteps, step)
			}
		}

		_pipeline.Steps = filteredSteps

		// check if any steps are left to run, excluding "init" step
		if len(_pipeline.Steps) <= 1 {
			return fmt.Errorf("no steps left to run after removing skipped steps")
		}
	}

	return nil
}
