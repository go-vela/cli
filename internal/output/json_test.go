// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"errors"
	"testing"
)

func TestOutput_JSON(t *testing.T) {
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
			input: map[string]string{"hello": "world"},
		},
		{ // slice
			failure: false,
			input: []interface{}{1, 2, 3},
		},
		{ // slice complex
			failure: false,
			input: []interface{}{struct{ Foo string }{Foo: "bar"}},
		},
		{ // complex
			failure: false,
			input: []struct{ Foo string }{{"bar"}, {"baz"}},
		},
		{
			failure: true,
			input:   new(failMarshaler),
		},
	}

	// run tests
	for _, test := range tests {
		err := JSON(test.input)

		if test.failure {
			if err == nil {
				t.Errorf("JSON should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("JSON returned err: %v", err)
		}
	}
}

func TestOutput_RawJSON(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		input   interface{}
	}{
		{
			failure: false,
			input:   "hello",
		},
		{
			failure: true,
			input:   new(failMarshaler),
		},
	}

	// run tests
	for _, test := range tests {
		err := RawJSON(test.input)

		if test.failure {
			if err == nil {
				t.Errorf("RawJSON should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("RawJSON returned err: %v", err)
		}
	}
}

type failMarshaler struct{}

func (f *failMarshaler) MarshalJSON() ([]byte, error) {
	return nil, errors.New("this is a marshaler that fails when you try to marshal")
}

func (f *failMarshaler) UnmarshalJSON([]byte) error {
	return errors.New("this is a marshaler that fails when you try to unmarshal")
}
