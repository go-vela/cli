// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
)

// RawJSON parses the provided input and
// renders the parsed input in raw JSON
// before outputting it to stdout.
func RawJSON(_input interface{}) error {
	// check if the input provided is nil
	if _input == nil {
		return errors.New("empty value provided for RawJSON output")
	}

	// check if the value of input provided is nil
	//
	// We are using reflect here due to the nature
	// of how interfaces work in Go. It is possible
	// for _input to be a non-nil interface but the
	// underlying value to be empty or nil.
	if reflect.ValueOf(_input).IsZero() {
		return errors.New("empty value provided for RawJSON output")
	}

	// marshal the input into raw JSON
	output, err := json.Marshal(_input)
	if err != nil {
		return err
	}

	// ensure we output to stdout
	fmt.Fprintln(os.Stdout, string(output))

	return nil
}
