// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"reflect"
	"testing"

	"github.com/go-vela/types/yaml"
)

func TestPipeline_steps(t *testing.T) {
	// setup tests
	tests := []struct {
		pipelineType string
		want         *yaml.Build
	}{
		{
			pipelineType: "",
			want:         steps(""),
		},
		{
			pipelineType: "go",
			want:         steps("go"),
		},
		{
			pipelineType: "node",
			want:         steps("node"),
		},
		{
			pipelineType: "java",
			want:         steps("java"),
		},
	}

	// run tests
	for _, test := range tests {
		got := steps(test.pipelineType)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("steps is %v, want %v", got, test.want)
		}
	}
}
