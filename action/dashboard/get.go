// SPDX-License-Identifier: Apache-2.0

package dashboard

import (
	"github.com/sirupsen/logrus"

	"github.com/go-vela/sdk-go/vela"
	api "github.com/go-vela/server/api/types"
)

// Get captures a list of dashboards based off the provided configuration.
func (c *Config) Get(client *vela.Client) error {
	logrus.Debug("executing get for dashboard configuration")

	// send API call to capture a list of dashboards
	dashCards, _, err := client.Dashboard.GetAllUser()
	if err != nil {
		return err
	}

	if c.Full {
		err = outputDashboard(dashCards, c)
	} else {
		dashboards := []*api.Dashboard{}

		for _, d := range *dashCards {
			dashboards = append(dashboards, d.Dashboard)
		}

		err = outputDashboard(dashboards, c)
	}

	return err
}
