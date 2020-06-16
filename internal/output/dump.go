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

// Dump outputs the provided input to stdout
// using github.com/davecgh/go-spew/spew to
// dump the input.
//
// More Information:
//
// Dump displays the passed parameters to standard
// out with newlines, customizable indentation,
// and additional debug information such as complete
// types and all pointer addresses used to indirect
// to the final value. It provides the following
// features over the built-in printing facilities
// provided by the fmt package:
//
// * Pointers are dereferenced and followed
// * Circular data structures are detected and
//   handled properly
// * Custom Stringer/error interfaces are
//   optionally invoked, including on
//   unexported types
// * Custom types which only implement the
//   Stringer/error interfaces via a pointer
//   receiver are optionally invoked when
//   passing non-pointer variables
// * Byte arrays and slices are dumped like
//   the hexdump -C command which includes
//   offsets, byte values in hex, and ASCII
//   output
func Dump(_input interface{}) error {
	// check if the input provided is nil
	if _input == nil {
		return errors.New("empty value provided for dump output")
	}

	// check if the value of input provided is nil
	//
	// We are using reflect here due to the nature
	// of how interfaces work in Go. It is possible
	// for _input to be a non-nil interface but the
	// underlying value to be empty or nil.
	if reflect.ValueOf(_input).IsZero() {
		return errors.New("empty value provided for dump output")
	}

	// ensure we output to stdout
	spew.Fdump(os.Stdout, _input)

	return nil
}
