// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/sdk-go/vela"
	api "github.com/go-vela/server/api/types"
)

// Add creates a dashboard based off the provided configuration.
func (c *Config) Add(ctx context.Context, client *vela.Client) error {
	logrus.Debug("executing add for dashboard configuration")

	dashRepos := []*api.DashboardRepo{}
	dashAdmins := []*api.User{}

	// generate dashboard repos
	for _, r := range c.AddRepos {
		repo := new(api.DashboardRepo)
		repo.SetName(r)

		if len(c.Branches) > 0 {
			repo.SetBranches(c.Branches)
		}

		if len(c.Events) > 0 {
			repo.SetEvents(c.Events)
		}

		dashRepos = append(dashRepos, repo)
	}

	// generate dashboard admins
	for _, u := range c.AddAdmins {
		admin := new(api.User)
		admin.SetName(u)

		dashAdmins = append(dashAdmins, admin)
	}

	// create the dashboard object
	d := &api.Dashboard{
		Name:   vela.String(c.Name),
		Repos:  &dashRepos,
		Admins: &dashAdmins,
	}

	logrus.Tracef("adding dashboard %s", c.Name)

	// send API call to add a dashboard
	dashboard, _, err := client.Dashboard.Add(ctx, d)
	if err != nil {
		return err
	}

	return outputDashboard(dashboard, c)
}
