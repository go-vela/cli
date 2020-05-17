// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// getAction defines the action for getting a list of builds.
const getAction = "get"

// helper function to capture the provided
// input and create the object used to
// capture a list of builds.
func get(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the build configuration
	b := &Build{
		Action:  getAction,
		Org:     c.String("org"),
		Repo:    c.String("repo"),
		Page:    c.Int("page"),
		PerPage: c.Int("per.page"),
		Output:  c.String("output"),
	}

	// validate build configuration
	err = b.Validate()
	if err != nil {
		return err
	}

	// execute the get call for the build configuration
	return b.Get(client)
}
