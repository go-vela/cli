// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"github.com/go-vela/sdk-go/vela"

	"github.com/urfave/cli/v2"
)

// viewAction defines the action for inspecting a build.
const viewAction = "view"

// helper function to capture the provided
// input and create the object used to
// inspect a build.
func view(c *cli.Context) error {
	// create a vela client
	client, err := vela.NewClient(c.String("addr"), nil)
	if err != nil {
		return err
	}

	// set token from global config
	client.Authentication.SetTokenAuth(c.String("token"))

	// create the build configuration
	b := &Build{
		Action: viewAction,
		Org:    c.String("org"),
		Repo:   c.String("repo"),
		Number: c.Int("build"),
		Output: c.String("output"),
	}

	// validate build configuration
	err = b.Validate()
	if err != nil {
		return err
	}

	// execute the view call for the build configuration
	return b.View(client)
}
