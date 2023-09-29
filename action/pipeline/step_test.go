// SPDX-License-Identifier: Apache-2.0

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
