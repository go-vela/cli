// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code among actions
package repo

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
	api "github.com/go-vela/server/api/types"
)

// Add creates a repository based off the provided configuration.
func (c *Config) Add(client *vela.Client) error {
	logrus.Debug("executing add for repo configuration")

	// create the repository object
	//
	// https://pkg.go.dev/github.com/go-vela/server/api/types?tab=doc#Repo
	r := &api.Repo{
		Org:          vela.String(c.Org),
		Name:         vela.String(c.Name),
		FullName:     vela.String(fmt.Sprintf("%s/%s", c.Org, c.Name)),
		Link:         vela.String(c.Link),
		Clone:        vela.String(c.Clone),
		Branch:       vela.String(c.Branch),
		BuildLimit:   vela.Int32(c.BuildLimit),
		Timeout:      vela.Int32(c.Timeout),
		Counter:      vela.Int64(c.Counter),
		Visibility:   vela.String(c.Visibility),
		Private:      vela.Bool(c.Private),
		Trusted:      vela.Bool(c.Trusted),
		Active:       vela.Bool(c.Active),
		PipelineType: vela.String(c.PipelineType),
		ApproveBuild: vela.String(c.ApproveBuild),
	}

	logrus.Tracef("adding repo %s/%s", c.Org, c.Name)

	if len(c.Events) > 0 {
		evs, err := api.NewEventsFromSlice(c.Events)
		if err != nil {
			return err
		}

		r.SetAllowEvents(evs)
	}

	// send API call to add a repository
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#RepoService.Add
	repo, _, err := client.Repo.Add(r)
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
		return output.JSON(repo, c.Color)
	case output.DriverSpew:
		// output the repository in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(repo)
	case output.DriverYAML:
		// output the repository in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(repo, c.Color)
	default:
		// output the repository in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(repo)
	}
}
