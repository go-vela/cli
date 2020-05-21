// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"os"

	"github.com/davecgh/go-spew/spew"
)

// Default outputs the provided input
// to stdout using the go-spew/spew
// package to pretty print the input.
func Default(_input interface{}) error {
	// ensure we output to stdout
	spew.Fprintf(os.Stdout, "%v\n", _input)

	return nil
}
