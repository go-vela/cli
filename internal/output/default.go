// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"errors"
	"os"
	"reflect"

	"github.com/davecgh/go-spew/spew"
)

// Default outputs the provided input
// to stdout using the go-spew/spew
// package to pretty print the input.
func Default(_input interface{}) error {
	// check if the input provided is nil
	if _input == nil {
		return errors.New("empty value provided for default output")
	}

	// check if the value of input provided is nil
	//
	// We are using reflect here due to the nature
	// of how interfaces work in Go. It is possible
	// for _input to be a non-nil interface but the
	// underlying value to be empty or nil.
	if reflect.ValueOf(_input).IsZero() {
		return errors.New("empty value provided for default output")
	}

	// ensure we output to stdout
	spew.Fprintf(os.Stdout, "%v\n", _input)

	return nil
}
