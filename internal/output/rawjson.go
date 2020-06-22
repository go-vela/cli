// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// RawJSON parses the provided input and
// renders the parsed input in raw JSON
// before outputting it to stdout.
func RawJSON(_input interface{}) error {
	logrus.Debugf("creating output with %s driver", DriverRawJSON)

	// validate the input provided
	err := validate(DriverRawJSON, _input)
	if err != nil {
		return err
	}

	// marshal the input into raw JSON
	output, err := json.Marshal(_input)
	if err != nil {
		return err
	}

	logrus.Tracef("sending output to stdout with %s driver", DriverRawJSON)

	// ensure we output to stdout
	fmt.Fprintln(os.Stdout, string(output))

	return nil
}
