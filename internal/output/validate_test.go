// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"testing"
)

func TestOutput_validate(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		driver  string
		input   interface{}
	}{
		{
			failure: false,
			driver:  DriverDefault,
			input:   "hello",
		},
		{ // map
			failure: false,
			driver:  DriverDefault,
			input:   map[string]string{"hello": "world"},
		},
		{ // slice
			failure: false,
			driver:  DriverDefault,
			input:   []interface{}{1, 2, 3},
		},
		{ // slice complex
			failure: false,
			driver:  DriverDefault,
			input:   []interface{}{struct{ Foo string }{Foo: "bar"}},
		},
		{ // complex
			failure: false,
			driver:  DriverDefault,
			input:   []struct{ Foo string }{{"bar"}, {"baz"}},
		},
		{
			failure: true,
			driver:  DriverDefault,
			input:   nil,
		},
		{
			failure: true,
			driver:  DriverDefault,
			input:   "",
		},
		{
			failure: true,
			driver:  DriverDump,
			input:   nil,
		},
		{
			failure: true,
			driver:  DriverDump,
			input:   "",
		},
		{
			failure: true,
			driver:  DriverJSON,
			input:   nil,
		},
		{
			failure: true,
			driver:  DriverJSON,
			input:   "",
		},
		{
			failure: true,
			driver:  DriverRawJSON,
			input:   nil,
		},
		{
			failure: true,
			driver:  DriverRawJSON,
			input:   "",
		},
		{
			failure: true,
			driver:  DriverSpew,
			input:   nil,
		},
		{
			failure: true,
			driver:  DriverSpew,
			input:   "",
		},
		{
			failure: true,
			driver:  DriverYAML,
			input:   nil,
		},
		{
			failure: true,
			driver:  DriverYAML,
			input:   "",
		},
	}

	// run tests
	for _, test := range tests {
		err := validate(test.driver, test.input)

		if test.failure {
			if err == nil {
				t.Errorf("validate should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("validate returned err: %v", err)
		}
	}
}
