// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// YAML parses the provided input and
// renders the parsed input in YAML
// before outputting it to stdout.
func YAML(_input interface{}) error {
	// marshal the input into YAML
	output, err := yaml.Marshal(_input)
	if err != nil {
		return err
	}

	// ensure we output to stdout
	fmt.Fprintln(os.Stdout, string(output))

	return nil
}
