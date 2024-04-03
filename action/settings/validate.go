// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"fmt"
	"net/url"

	"github.com/go-vela/cli/internal"
	"github.com/sirupsen/logrus"
)

// Validate verifies the configuration provided.
func (c *Config) Validate() error {
	logrus.Debug("validating worker configuration")

	// address is required for adding a worker, validate it
	if c.Action == internal.ActionAdd {
		if len(c.Address) == 0 {
			return fmt.Errorf("no worker address provided")
		}

		_, err := url.Parse(c.Address)
		if err != nil {
			return fmt.Errorf("error while parsing worker address provided")
		}
	}

	// address is optional for update, but validate it if provided
	if c.Action == internal.ActionUpdate && len(c.Address) > 0 {
		_, err := url.Parse(c.Address)
		if err != nil {
			return fmt.Errorf("error while parsing worker address provided")
		}
	}

	// anything other than "get" and "add" action needs hostname
	if (c.Action != internal.ActionGet && c.Action != internal.ActionAdd) && len(c.Hostname) == 0 {
		return fmt.Errorf("no worker hostname provided")
	}

	return nil
}
