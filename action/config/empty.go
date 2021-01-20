// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

import (
	"github.com/sirupsen/logrus"
)

// Empty checks if the config file contains empty values.
func (c *ConfigFile) Empty() bool {
	logrus.Debug("checking if config file is empty")

	// check if the API object is nil
	if c.API != nil {
		logrus.Trace("checking if API values in config file are empty")

		// check if the API address is set
		if len(c.API.Address) > 0 {
			return false
		}

		// check if the API token is set
		if len(c.API.Token) > 0 {
			return false
		}

		// check if the API access token is set
		if len(c.API.AccessToken) > 0 {
			return false
		}

		// check if the API refresh token is set
		if len(c.API.RefreshToken) > 0 {
			return false
		}

		// check if the API version is set
		if len(c.API.Version) > 0 {
			return false
		}
	}

	// check if the log object is nil
	if c.Log != nil {
		logrus.Trace("checking if log values in config file are empty")

		// check if the log level is set
		if len(c.Log.Level) > 0 {
			return false
		}
	}

	// check if the secret object is nil
	if c.Secret != nil {
		logrus.Trace("checking if secret values in config file are empty")

		// check if the secret engine is set
		if len(c.Secret.Engine) > 0 {
			return false
		}

		// check if the secret type is set
		if len(c.Secret.Type) > 0 {
			return false
		}
	}

	// check if the output is set
	if len(c.Output) > 0 {
		return false
	}

	// check if the org is set
	if len(c.Org) > 0 {
		return false
	}

	// check if the repo is set
	if len(c.Repo) > 0 {
		return false
	}

	return true
}
