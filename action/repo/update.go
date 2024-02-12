// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code among actions
package repo

import (
	"fmt"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/go-vela/types/library"

	"github.com/sirupsen/logrus"
)

// Update modifies a repository based off the provided configuration.
func (c *Config) Update(client *vela.Client) error {
	logrus.Debug("executing update for repo configuration")

	// create the repository object
	//
	// https://pkg.go.dev/github.com/go-vela/types/library?tab=doc#Repo
	r := &library.Repo{
		Org:          vela.String(c.Org),
		Name:         vela.String(c.Name),
		FullName:     vela.String(fmt.Sprintf("%s/%s", c.Org, c.Name)),
		Link:         vela.String(c.Link),
		Clone:        vela.String(c.Clone),
		Branch:       vela.String(c.Branch),
		BuildLimit:   vela.Int64(c.BuildLimit),
		Timeout:      vela.Int64(c.Timeout),
		Counter:      vela.Int(c.Counter),
		Visibility:   vela.String(c.Visibility),
		Private:      vela.Bool(c.Private),
		Trusted:      vela.Bool(c.Trusted),
		Active:       vela.Bool(c.Active),
		PipelineType: vela.String(c.PipelineType),
		ApproveBuild: vela.String(c.ApproveBuild),
	}

	if len(c.Events) > 0 {
		populateEvents(r, c.Events)
	}

	logrus.Tracef("updating repo %s/%s", c.Org, c.Name)

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
