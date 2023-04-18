// Copyright (c) 2023 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Stderr outputs the provided input to stderr.
func Stderr(_input interface{}) error {
	logrus.Debugf("creating output with %s driver", DriverStderr)

	// validate the input provided
	err := validate(DriverStderr, _input)
	if err != nil {
		return err
	}

	logrus.Tracef("sending output to stderr with %s driver", DriverStderr)

	// ensure we output to stderr
	fmt.Fprintln(os.Stderr, _input)

	return nil
}
