// SPDX-License-Identifier: Apache-2.0

//nolint:dupl // ignore similar code among actions
package repo

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/sdk-go/vela"
	api "github.com/go-vela/server/api/types"
)

// Add creates a repository based off the provided configuration.
func (c *Config) Add(ctx context.Context, client *vela.Client) error {
	logrus.Debug("executing add for repo configuration")

	// create the repository object
	//
	// https://pkg.go.dev/github.com/go-vela/server/api/types?tab=doc#Repo
	r := &api.Repo{
		Org:              new(c.Org),
		Name:             new(c.Name),
		FullName:         new(fmt.Sprintf("%s/%s", c.Org, c.Name)),
		Link:             new(c.Link),
		Clone:            new(c.Clone),
		Branch:           new(c.Branch),
		BuildLimit:       new(c.BuildLimit),
		Timeout:          new(c.Timeout),
		Counter:          new(c.Counter),
		Visibility:       new(c.Visibility),
		Private:          new(c.Private),
		Trusted:          new(c.Trusted),
		Active:           new(c.Active),
		PipelineType:     new(c.PipelineType),
		ApproveBuild:     new(c.ApproveBuild),
		MergeQueueEvents: new(c.MergeQueueEvents),
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
	repo, _, err := client.Repo.Add(ctx, r)
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
