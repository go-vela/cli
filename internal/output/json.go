// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSON parses the provided input and
// renders the parsed input in pretty
// JSON before outputting it to stdout.
func JSON(_input interface{}) error {
	// marshal the input into pretty JSON
	output, err := json.MarshalIndent(_input, "", "    ")
	if err != nil {
		return err
	}

	// ensure we output to stdout
	fmt.Fprintln(os.Stdout, string(output))

	return nil
}

// RawJSON parses the provided input and
// renders the parsed input in raw JSON
// before outputting it to stdout.
func RawJSON(_input interface{}) error {
	// marshal the input into raw JSON
	output, err := json.Marshal(_input)
	if err != nil {
		return err
	}

	// ensure we output to stdout
	fmt.Fprintln(os.Stdout, string(output))

	return nil
}
