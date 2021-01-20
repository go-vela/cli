// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"testing"
)

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
