// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

import (
	"io/ioutil"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/sdk-go/vela"

	"github.com/sirupsen/logrus"
)

// Cancel cancels a build based off the provided configuration.
func (c *Config) Cancel(client *vela.Client) error {
	logrus.Debug("executing cancel for build configuration")

	logrus.Tracef("canceling build %s/%s/%d", c.Org, c.Repo, c.Number)

	// send API call to cancel a build
	//
	// https://pkg.go.dev/github.com/go-vela/sdk-go/vela?tab=doc#BuildService.Cancel
	_, resp, err := client.Build.Cancel(c.Org, c.Repo, c.Number)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	respS := string(respBody)

	// handle the output based off the provided configuration
	switch c.Output {
	case output.DriverDump:
		// output the build in dump format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Dump
		return output.Dump(respS)
	case output.DriverJSON:
		// output the build in JSON format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#JSON
		return output.JSON(respS)
	case output.DriverSpew:
		// output the build in spew format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Spew
		return output.Spew(respS)
	case output.DriverYAML:
		// output the build in YAML format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#YAML
		return output.YAML(&respS)
	default:
		// output the build in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(respS)
	}
}
