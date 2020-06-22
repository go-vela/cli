// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"os"

	"github.com/davecgh/go-spew/spew"

	"github.com/sirupsen/logrus"
)

// Spew outputs the provided input to stdout
// using github.com/davecgh/go-spew/spew to
// verbosely print the input.
func Spew(_input interface{}) error {
	logrus.Debugf("creating output with %s driver", DriverSpew)

	// validate the input provided
	err := validate(DriverSpew, _input)
	if err != nil {
		return err
	}

	logrus.Tracef("sending output to stdout with %s driver", DriverSpew)

	// ensure we output to stdout
	spew.Fprintf(os.Stdout, "%#+v\n", _input)

	return nil
}
