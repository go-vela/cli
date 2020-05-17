// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"errors"
	"testing"
)

func TestOutput_YAML(t *testing.T) {
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
		err := YAML(test.input)

		if test.failure {
			if err == nil {
				t.Errorf("YAML should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("YAML returned err: %v", err)
		}
	}
}

func (f *failMarshaler) MarshalYAML() (interface{}, error) {
	return nil, errors.New("this is a marshaler that fails when you try to marshal")
}

func (f *failMarshaler) UnmarshalYAML(unmarshal func(interface{}) error) error {
	return errors.New("this is a marshaler that fails when you try to unmarshal")
}
