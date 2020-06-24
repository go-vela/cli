// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Stdout outputs the provided input to stdout.
func Stdout(_input interface{}) error {
	logrus.Debugf("creating output with %s driver", DriverStdout)

	// validate the input provided
	err := validate(DriverStdout, _input)
	if err != nil {
		return err
	}

	logrus.Tracef("sending output to stdout with %s driver", DriverStdout)

	// ensure we output to stdout
	fmt.Fprintln(os.Stdout, _input)

	return nil
}
