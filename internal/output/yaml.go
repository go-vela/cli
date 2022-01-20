// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/sirupsen/logrus"
)

// YAML parses the provided input and
// renders the parsed input in YAML
// before outputting it to stdout.
func YAML(_input interface{}) error {
	logrus.Debugf("creating output with %s driver", DriverYAML)

	// validate the input provided
	err := validate(DriverYAML, _input)
	if err != nil {
		return err
	}

	// marshal the input into YAML
	output, err := yaml.Marshal(_input)
	if err != nil {
		return err
	}

	logrus.Tracef("sending output to stdout with %s driver", DriverYAML)

	// ensure we output to stdout
	fmt.Fprintln(os.Stdout, string(output))

	return nil
}
