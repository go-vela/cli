// SPDX-License-Identifier: Apache-2.0

package user

import (
	"context"
	"fmt"
	"slices"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/sdk-go/vela"
	api "github.com/go-vela/server/api/types"
)

// Update modifies a dashboard based off the provided configuration.
func (c *Config) Update(ctx context.Context, client *vela.Client) error {
	logrus.Debug("executing update for dashboard configuration")

	var (
		user *api.User
		err  error
	)

	if len(c.Name) > 0 {
		user, _, err = client.User.Get(ctx, c.Name)
	} else {
		user, _, err = client.User.GetCurrent(ctx)
	}

	if err != nil {
		return err
	}

	// drop specified dashboards from the user
	if len(c.DropDashboards) > 0 {
		newDashboards := []string{}

		for _, d := range user.GetDashboards() {
			if !slices.Contains(c.DropDashboards, d) {
				newDashboards = append(newDashboards, d)
			}
		}

		user.SetDashboards(newDashboards)
	}

	// add specified repositories to the dashboard
	if len(c.AddDashboards) > 0 {
		dashboards := user.GetDashboards()

		for _, d := range c.AddDashboards {
			_, _, err := client.Dashboard.Get(ctx, d)
			if err != nil {
				return fmt.Errorf("unable to get dashboard %s: %w", d, err)
			}

			dashboards = append(dashboards, d)
		}

		user.SetDashboards(dashboards)
	}

	// drop specified favorites from the user
	if len(c.DropFavorites) > 0 {
		newFavorites := []string{}

		for _, f := range user.GetFavorites() {
			if !slices.Contains(c.DropFavorites, f) {
				newFavorites = append(newFavorites, f)
			}
		}

		user.SetFavorites(newFavorites)
	}

	// add specified favorites to the user
	if len(c.AddFavorites) > 0 {
		favorites := user.GetFavorites()

		for _, f := range c.AddFavorites {
			splitRepo := strings.Split(f, "/")

			if len(splitRepo) != 2 {
				return fmt.Errorf("invalid format for repository: %s (valid format: <org>/<repo>)", f)
			}

			_, _, err := client.Repo.Get(ctx, splitRepo[0], splitRepo[1])
			if err != nil {
				return fmt.Errorf("unable to get repo %s: %w", f, err)
			}

			favorites = append(favorites, f)
		}

		user.SetFavorites(favorites)
	}

	// send API call to modify the user
	if len(c.Name) > 0 {
		user, _, err = client.User.Update(ctx, c.Name, user)
	} else {
		user, _, err = client.User.UpdateCurrent(ctx, user)
	}

	if err != nil {
		return err
	}

	return outputUser(user, c)
}
