// SPDX-License-Identifier: Apache-2.0

package user

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/sdk-go/vela"
	api "github.com/go-vela/server/api/types"
)

// View inspects a user based off the provided configuration.
func (c *Config) View(ctx context.Context, client *vela.Client) error {
	logrus.Debug("executing view for user configuration")

	var (
		user *api.User
		err  error
	)

	// send API call to capture user
	if len(c.Name) > 0 {
		user, _, err = client.User.Get(ctx, c.Name)
	} else {
		user, _, err = client.User.GetCurrent(ctx)
	}

	if err != nil {
		return err
	}

	return outputUser(user, c)
}
