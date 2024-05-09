// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"fmt"
	"slices"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/sdk-go/vela"
	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/constants"
)

// Update modifies a dashboard based off the provided configuration.
func (c *Config) Update(client *vela.Client) error {
	logrus.Debug("executing update for dashboard configuration")

	dashCard, _, err := client.Dashboard.Get(c.ID)
	if err != nil {
		return err
	}

	// pull dashboard metadata from the API response
	dashboard := dashCard.Dashboard

	// drop specified repositories from the dashboard
	if len(c.DropRepos) > 0 {
		newRepos := []*api.DashboardRepo{}

		for _, r := range dashboard.GetRepos() {
			if !slices.Contains(c.DropRepos, r.GetName()) {
				newRepos = append(newRepos, r)
			}
		}

		dashboard.SetRepos(newRepos)
	}

	// add specified repositories from the dashboard
	if len(c.AddRepos) > 0 {
		repos := dashboard.GetRepos()

		for _, r := range c.AddRepos {
			repo := new(api.DashboardRepo)
			repo.SetName(r)

			if len(c.Branches) > 0 {
				repo.SetBranches(c.Branches)
			}

			if len(c.Events) > 0 {
				repo.SetEvents(c.Events)
			}

			repos = append(repos, repo)
		}

		dashboard.SetRepos(repos)
	}

	// update specified repositories from the dashboard
	if len(c.TargetRepos) > 0 {
		repos := dashboard.GetRepos()
		for _, r := range repos {
			if slices.Contains(c.TargetRepos, r.GetName()) {
				if len(c.Branches) > 0 {
					r.SetBranches(c.Branches)
				}

				if len(c.Events) > 0 {
					r.SetEvents(c.Events)
				}
			}
		}

		dashboard.SetRepos(repos)
	}

	// drop specified admins from the dashboard
	if len(c.DropAdmins) > 0 {
		newAdmins := []*api.User{}

		for _, a := range dashboard.GetAdmins() {
			if !slices.Contains(c.DropAdmins, a.GetName()) {
				newAdmins = append(newAdmins, a)
			}
		}

		dashboard.SetAdmins(newAdmins)
	}

	// add specified admins from the dashboard
	if len(c.AddAdmins) > 0 {
		admins := dashboard.GetAdmins()

		for _, a := range c.AddAdmins {
			admin := new(api.User)
			admin.SetName(a)

			admins = append(admins, admin)
		}

		dashboard.SetAdmins(admins)
	}

	// update the name of the dashboard
	if len(c.Name) > 0 {
		dashboard.SetName(c.Name)
	}

	// verify the number of repositories for a dashboard
	if len(dashboard.GetRepos()) > constants.DashboardRepoLimit {
		return fmt.Errorf("maximum number of repositories for a dashboard is %d", constants.DashboardRepoLimit)
	}

	// send API call to modify a dashboard
	dashboard, _, err = client.Dashboard.Update(dashboard)
	if err != nil {
		return err
	}

	return outputDashboard(dashboard, c)
}
