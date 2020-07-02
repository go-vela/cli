// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"fmt"
	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"
)

// Update modifies a repository based off the provided configuration.
func (c *Config) Update(client *vela.Client) error {
	// create the repository object
	//
	// https://pkg.go.dev/github.com/go-vela/types/library?tab=doc#Repo
	r := &library.Repo{
		Org:        vela.String(c.Org),
		Name:       vela.String(c.Name),
		FullName:   vela.String(fmt.Sprintf("%s/%s", c.Org, c.Name)),
		Link:       vela.String(c.Link),
		Clone:      vela.String(c.Clone),
		Branch:     vela.String(c.Branch),
		Timeout:    vela.Int64(c.Timeout),
		Visibility: vela.String(c.Visibility),
		Private:    vela.Bool(c.Private),
		Trusted:    vela.Bool(c.Trusted),
		Active:     vela.Bool(c.Active),
	}

	// iterate through all events provided
	for _, event := range c.Events {
		// check if the repository should allow push events
		if event == constants.EventPush {
			r.AllowPush = vela.Bool(true)
		}

		// check if the repository should allow pull_request events
		if event == constants.EventPull {
			r.AllowPull = vela.Bool(true)
		}

		// check if the repository should allow tag events
		if event == constants.EventTag {
			r.AllowTag = vela.Bool(true)
		}

		// check if the repository should allow deployment events
		if event == constants.EventDeploy {
			r.AllowDeploy = vela.Bool(true)
		}

		// check if the repository should allow comment events
		if event == constants.EventComment {
			r.AllowComment = vela.Bool(true)
		}
	}

	// send API call to modify a repository
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#RepoService.Update
	repo, _, err := client.Repo.Update(c.Org, c.Name, r)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the repository in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(repo)
	case output.DriverJSON:
		// output the repository in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(repo)
	case output.DriverSpew:
		// output the repository in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(repo)
	case output.DriverYAML:
		// output the repository in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(repo)
	default:
		// output the repository in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(repo)
	}
}
