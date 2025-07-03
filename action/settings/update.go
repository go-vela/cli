// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"slices"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

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
			CloneImage:        c.CloneImage,
			TemplateDepth:     c.TemplateDepth,
			StarlarkExecLimit: c.StarlarkExecLimit,
		},
		SCM: &settings.SCM{
			RepoRoleMap: c.RepoRoleMap,
			OrgRoleMap:  c.OrgRoleMap,
			TeamRoleMap: c.TeamRoleMap,
		},
		RepoAllowlist:     vela.Strings(s.GetRepoAllowlist()),
		ScheduleAllowlist: vela.Strings(s.GetScheduleAllowlist()),
	}

	// update max dashboard repos if set
	if c.MaxDashboardRepos > 0 && c.MaxDashboardRepos != s.GetMaxDashboardRepos() {
		sUpdate.SetMaxDashboardRepos(c.MaxDashboardRepos)
	}

	// drop specified routes
	if len(c.DropRoutes) > 0 {
		newRoutes := []string{}

		for _, r := range sUpdate.GetRoutes() {
			if !slices.Contains(c.DropRoutes, r) {
				newRoutes = append(newRoutes, r)
			}
		}

		sUpdate.SetRoutes(newRoutes)
	}

	// add specified routes
	if len(c.AddRoutes) > 0 {
		routes := sUpdate.GetRoutes()

		for _, r := range c.AddRoutes {
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

	if c.Routes != nil {
		sUpdate.Routes = c.Routes
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

		if f.RepoAllowlist != nil {
			s.RepoAllowlist = f.RepoAllowlist
		}

		if f.ScheduleAllowlist != nil {
			s.ScheduleAllowlist = f.ScheduleAllowlist
		}

		// update values if set
		if f.Compiler != nil {
			if f.CloneImage != nil {
				s.CloneImage = vela.String(f.GetCloneImage())
			}

			if f.TemplateDepth != nil {
				s.TemplateDepth = vela.Int(f.GetTemplateDepth())
			}

			if f.StarlarkExecLimit != nil {
				s.StarlarkExecLimit = vela.Int64(f.GetStarlarkExecLimit())
			}
		}

		if f.Queue != nil {
			if f.Routes != nil {
				s.Routes = f.Routes
			}
		}

		if f.SCM != nil {
			if f.RepoRoleMap != nil {
				s.RepoRoleMap = f.RepoRoleMap
			}

			if f.OrgRoleMap != nil {
				s.OrgRoleMap = f.OrgRoleMap
			}

			if f.TeamRoleMap != nil {
				s.TeamRoleMap = f.TeamRoleMap
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
