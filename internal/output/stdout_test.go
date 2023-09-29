// SPDX-License-Identifier: Apache-2.0

package output

import (
	"testing"
)

func TestOutput_Default(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		input   interface{}
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
		err := Stdout(test.input)

		if test.failure {
			if err == nil {
				t.Errorf("Stdout should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Stdout returned err: %v", err)
		}
	}
}
