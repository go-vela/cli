// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"testing"
)

func TestOutput_Dump(t *testing.T) {
	// setup tests
	tests := []struct {
		input   interface{}
		failure bool
	}{
		{
			failure: false,
			input:   "hello",
		},
		{ // map
			failure: false,
			input:   map[string]string{"hello": "world"},
		},
		{ // slice
			failure: false,
			input:   []interface{}{1, 2, 3},
		},
		{ // slice complex
			failure: false,
			input:   []interface{}{struct{ Foo string }{Foo: "bar"}},
		},
		{ // complex
			failure: false,
			input:   []struct{ Foo string }{{"bar"}, {"baz"}},
		},
		{
			failure: true,
			input:   nil,
		},
		{
			failure: true,
			input:   "",
		},
	}

	// run tests
	for _, test := range tests {
		err := Dump(test.input)

		if test.failure {
			if err == nil {
				t.Errorf("Dump should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Dump returned err: %v", err)
		}
	}
}
