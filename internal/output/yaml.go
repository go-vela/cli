// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"gopkg.in/yaml.v2"

	"github.com/sirupsen/logrus"
)

// YAML parses the provided input and
// renders the parsed input in YAML
// before outputting it to stdout.
func YAML(_input interface{}) error {
	logrus.Debugf("creating output with %s driver", DriverYAML)

	// check if the input provided is nil
	if _input == nil {
		return errors.New("empty value provided for YAML output")
	}

	// check if the value of input provided is nil
	//
	// We are using reflect here due to the nature
	// of how interfaces work in Go. It is possible
	// for _input to be a non-nil interface but the
	// underlying value to be empty or nil.
	if reflect.ValueOf(_input).IsZero() {
		return errors.New("empty value provided for YAML output")
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
