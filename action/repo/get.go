// SPDX-License-Identifier: Apache-2.0

package repo

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
)

// Get captures a list of repositories based off the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	logrus.Debug("executing get for repo configuration")

	// set the pagination options for list of repositories
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#ListOptions
	opts := &vela.ListOptions{
		Page:    c.Page,
		PerPage: c.PerPage,
	}

	logrus.Tracef("capturing repos for current user")

	// send API call to capture a list of repositories
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#RepoService.GetAll
	repos, _, err := client.Repo.GetAll(opts)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the repositories in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(repos)
	case output.DriverJSON:
		// output the repositories in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(repos, c.Color)
	case output.DriverSpew:
		// output the repositories in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(repos)
	case "wide":
		// output the repositories in wide table format
		return wideTable(repos)
	case output.DriverYAML:
		// output the repositories in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(repos, c.Color)
	default:
		// output the repositories in table format
		return table(repos)
	}
}
