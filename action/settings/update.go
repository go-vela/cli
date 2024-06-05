// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"slices"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	"github.com/go-vela/cli/internal"
	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/api/types/settings"
)

// Update modifies settings based off the provided configuration.
func (c *Config) Update(client *vela.Client) error {
	logrus.Debug("executing update for settings configuration")

	// send API call to retrieve current settings
	s, _, err := client.Admin.Settings.Get()
	if err != nil {
		return err
	}

	// create the settings object
	sUpdate := &settings.Platform{
		Queue: &settings.Queue{
			Routes: vela.Strings(s.GetRoutes()),
		},
		Compiler: &settings.Compiler{
			CloneImage:        c.Compiler.CloneImage,
			TemplateDepth:     c.Compiler.TemplateDepth,
			StarlarkExecLimit: c.Compiler.StarlarkExecLimit,
		},
		RepoAllowlist:     vela.Strings(s.GetRepoAllowlist()),
		ScheduleAllowlist: vela.Strings(s.GetScheduleAllowlist()),
	}

	// drop specified routes
	if len(c.Queue.DropRoutes) > 0 {
		newRoutes := []string{}

		for _, r := range sUpdate.GetRoutes() {
			if !slices.Contains(c.Queue.DropRoutes, r) {
				newRoutes = append(newRoutes, r)
			}
		}

		sUpdate.SetRoutes(newRoutes)
	}

	// add specified routes
	if len(c.Queue.AddRoutes) > 0 {
		routes := sUpdate.GetRoutes()

		for _, r := range c.Queue.AddRoutes {
			if !slices.Contains(routes, r) {
				routes = append(routes, r)
			}
		}

		sUpdate.SetRoutes(routes)
	}

	// drop specified repositories from the allowlist
	if len(c.RepoAllowlistDropRepos) > 0 {
		newRepos := []string{}

		for _, r := range sUpdate.GetRepoAllowlist() {
			if !slices.Contains(c.RepoAllowlistDropRepos, r) {
				newRepos = append(newRepos, r)
			}
		}

		sUpdate.SetRepoAllowlist(newRepos)
	}

	// add specified repositories to the allowlist
	if len(c.RepoAllowlistAddRepos) > 0 {
		repos := sUpdate.GetRepoAllowlist()

		for _, r := range c.RepoAllowlistAddRepos {
			if !slices.Contains(repos, r) {
				repos = append(repos, r)
			}
		}

		sUpdate.SetRepoAllowlist(repos)
	}

	// drop specified repositories from the allowlist
	if len(c.ScheduleAllowlistDropRepos) > 0 {
		newRepos := []string{}

		for _, r := range sUpdate.GetScheduleAllowlist() {
			if !slices.Contains(c.ScheduleAllowlistDropRepos, r) {
				newRepos = append(newRepos, r)
			}
		}

		sUpdate.SetScheduleAllowlist(newRepos)
	}

	// add specified repositories to the allowlist
	if len(c.ScheduleAllowlistAddRepos) > 0 {
		repos := sUpdate.GetScheduleAllowlist()

		for _, r := range c.ScheduleAllowlistAddRepos {
			if !slices.Contains(repos, r) {
				repos = append(repos, r)
			}
		}

		sUpdate.SetScheduleAllowlist(repos)
	}

	// manual overrides (from file)
	if c.RepoAllowlist != nil {
		sUpdate.RepoAllowlist = c.RepoAllowlist
	}

	if c.ScheduleAllowlist != nil {
		sUpdate.ScheduleAllowlist = c.ScheduleAllowlist
	}

	if c.Queue.Routes != nil {
		sUpdate.Queue.Routes = c.Queue.Routes
	}

	logrus.Trace("updating settings")

	// send API call to modify settings
	sUpdated, _, err := client.Admin.Settings.Update(sUpdate)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(sUpdated)
	case output.DriverJSON:
		// output in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(sUpdated, c.Color)
	case output.DriverSpew:
		// output in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(sUpdated)
	case output.DriverYAML:
		// output in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(sUpdated, c.Color)
	default:
		// output in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(sUpdated)
	}
}

// UpdateFromFile updates from a file based on the provided configuration.
func (c *Config) UpdateFromFile(client *vela.Client) error {
	logrus.Debug("executing update from file for platform settings configuration")

	// capture absolute path to file
	path, err := filepath.Abs(c.File)
	if err != nil {
		return err
	}

	logrus.Tracef("reading platform settings contents from %s", path)

	// read contents of file
	contents, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	input := yaml.NewDecoder(bytes.NewReader(contents))

	f := new(ConfigFile)

	for input.Decode(f) == nil {
		if f.Platform == nil {
			return errors.New("invalid input, expected key 'platform'")
		}

		s := &Config{
			Action:   internal.ActionUpdate,
			Output:   c.Output,
			Compiler: Compiler{},
			Queue:    Queue{},
		}

		if f.Platform.RepoAllowlist != nil {
			s.RepoAllowlist = f.Platform.RepoAllowlist
		}

		if f.Platform.ScheduleAllowlist != nil {
			s.ScheduleAllowlist = f.Platform.ScheduleAllowlist
		}

		// update values if set
		if f.Compiler != nil {
			if f.Compiler.CloneImage != nil {
				s.Compiler.CloneImage = vela.String(f.Compiler.GetCloneImage())
			}

			if f.Compiler.TemplateDepth != nil {
				s.Compiler.TemplateDepth = vela.Int(f.Compiler.GetTemplateDepth())
			}

			if f.Compiler.StarlarkExecLimit != nil {
				s.Compiler.StarlarkExecLimit = vela.UInt64(f.Compiler.GetStarlarkExecLimit())
			}
		}

		if f.Queue != nil {
			if f.Queue.Routes != nil {
				s.Queue.Routes = f.Queue.Routes
			}
		}

		err = s.Validate()
		if err != nil {
			return err
		}

		err = s.Update(client)
		if err != nil {
			return err
		}
	}

	return nil
}
