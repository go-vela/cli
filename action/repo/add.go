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

// Add creates a repository based off the provided configuration.
func (c *Config) Add(client *vela.Client) error {
	// create the repository object
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

	for _, event := range c.Events {
		if event == constants.EventPush {
			r.AllowPush = vela.Bool(true)
		}

		if event == constants.EventPull {
			r.AllowPull = vela.Bool(true)
		}

		if event == constants.EventTag {
			r.AllowTag = vela.Bool(true)
		}

		if event == constants.EventDeploy {
			r.AllowDeploy = vela.Bool(true)
		}

		if event == constants.EventComment {
			r.AllowComment = vela.Bool(true)
		}
	}

	// send API call to add a repository
	repo, _, err := client.Repo.Add(r)
	if err != nil {
		return err
	}

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the repository in dump format
		err := output.Dump(repo)
		if err != nil {
			return err
		}
	case output.DriverJSON:
		// output the repository in JSON format
		err := output.JSON(repo)
		if err != nil {
			return err
		}
	case output.DriverSpew:
		// output the repository in spew format
		err := output.Spew(repo)
		if err != nil {
			return err
		}
	case output.DriverYAML:
		// output the repository in YAML format
		err := output.YAML(repo)
		if err != nil {
			return err
		}
	default:
		// output the repository in default format
		err := output.Default(repo)
		if err != nil {
			return err
		}
	}

	return nil
}
