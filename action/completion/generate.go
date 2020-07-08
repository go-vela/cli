// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package completion

import (
	"fmt"

	"github.com/go-vela/cli/internal/output"

	"github.com/sirupsen/logrus"
)

// Generate produces a script used to enable
// automatic completion for a Unix shell
// based off the provided configuration.
func (c *Config) Generate() error {
	logrus.Debug("executing generate for completion configuration")

	// generate the script based off the provided configuration
	switch {
	case c.Bash:
		logrus.Tracef("creating bash automatic completion script")

		// output the bash auto completion in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(BashAutoComplete)
	case c.Zsh:
		logrus.Tracef("creating zsh automatic completion script")

		// output the zsh auto completion in stdout format
		//
		// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
		return output.Stdout(ZshAutoComplete)
	default:
		// produce an invalid shell error by default
		return fmt.Errorf("invalid shell provided for completion")
	}
}
