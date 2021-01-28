// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

import (
	"reflect"
	"testing"

	"github.com/go-vela/types/yaml"
)

func TestPipeline_stages(t *testing.T) {
	// setup tests
	tests := []struct {
		pipelineType string
		want         *yaml.Build
	}{
		{
			pipelineType: "",
			want:         stages(""),
		},
		{
			pipelineType: "go",
			want:         stages("go"),
		},
		{
			pipelineType: "node",
			want:         stages("node"),
		},
		{
			pipelineType: "java",
			want:         stages("java"),
		},
	}

	// run tests
	for _, test := range tests {
		got := stages(test.pipelineType)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("stages is %v, want %v", got, test.want)
		}
	}
}
