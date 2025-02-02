// SPDX-License-Identifier: Apache-2.0

package pipeline

import (
	"io"
	"reflect"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/go-vela/server/compiler/types/pipeline"
)

func TestCollectMissingSecrets(t *testing.T) {
	// Set up test cases
	tests := []struct {
		name     string
		pipeline *pipeline.Build
		envVars  map[string]string
		want     map[string]string
	}{
		{
			name:     "pipeline nil",
			pipeline: nil,
			want:     map[string]string{},
		},
		{
			name: "no secrets",
			pipeline: &pipeline.Build{
				Steps: []*pipeline.Container{
					{Name: "step1"},
				},
			},
			want: map[string]string{},
		},
		{
			name: "missing step secret",
			pipeline: &pipeline.Build{
				Steps: []*pipeline.Container{
					{
						Name: "step1",
						Secrets: pipeline.StepSecretSlice{
							{
								Source: "source",
								Target: "TARGET_SECRET",
							},
						},
					},
				},
			},
			want: map[string]string{"[step: step1]": "TARGET_SECRET"},
		},
		{
			name: "provided step secret",
			pipeline: &pipeline.Build{
				Steps: []*pipeline.Container{
					{
						Name: "step1",
						Secrets: pipeline.StepSecretSlice{
							{
								Source: "source",
								Target: "TARGET_SECRET",
							},
						},
					},
				},
			},
			envVars: map[string]string{"TARGET_SECRET": "value"},
			want:    map[string]string{},
		},
		{
			name: "stage with missing secret",
			pipeline: &pipeline.Build{
				Stages: []*pipeline.Stage{
					{
						Name: "stage1",
						Steps: []*pipeline.Container{
							{
								Name: "step1",
								Secrets: pipeline.StepSecretSlice{
									{
										Source: "source",
										Target: "STAGE_SECRET",
									},
								},
							},
						},
					},
				},
			},
			want: map[string]string{"[stage: stage1][step: step1]": "STAGE_SECRET"},
		},
		{
			name: "provided step secret but value empty",
			pipeline: &pipeline.Build{
				Steps: []*pipeline.Container{
					{
						Name: "step1",
						Secrets: pipeline.StepSecretSlice{
							{
								Source: "source",
								Target: "TARGET_SECRET",
							},
						},
					},
				},
			},
			envVars: map[string]string{"TARGET_SECRET": ""},
			want:    map[string]string{},
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up environment
			for k, v := range tt.envVars {
				t.Setenv(k, v)
			}

			got := collectMissingSecrets(tt.pipeline)
			if len(got) != len(tt.want) {
				t.Errorf("collectMissingSecrets() = %v, want %v", got, tt.want)
			}

			for k, v := range tt.want {
				if got[k] != v {
					t.Errorf("collectMissingSecrets()[%s] = %v, want %v", k, got[k], v)
				}
			}
		})
	}
}

func TestSkipSteps(t *testing.T) {
	tests := []struct {
		name          string
		pipeline      *pipeline.Build
		stepsToRemove []string
		wantSteps     []string
		wantErr       bool
	}{
		{
			name: "skip one step",
			pipeline: &pipeline.Build{
				Steps: []*pipeline.Container{
					{Name: "init"},
					{Name: "step1"},
					{Name: "step2"},
				},
			},
			stepsToRemove: []string{"step1"},
			wantSteps:     []string{"init", "step2"},
			wantErr:       false,
		},
		{
			name: "skip all steps except init",
			pipeline: &pipeline.Build{
				Steps: []*pipeline.Container{
					{Name: "init"},
					{Name: "step1"},
				},
			},
			stepsToRemove: []string{"step1"},
			wantSteps:     []string{"init"},
			wantErr:       true,
		},
		{
			name: "skip steps in stages",
			pipeline: &pipeline.Build{
				Stages: []*pipeline.Stage{
					{
						Name: "stage1",
						Steps: []*pipeline.Container{
							{Name: "init"},
							{Name: "step1"},
							{Name: "step2"},
						},
					},
				},
			},
			stepsToRemove: []string{"step1"},
			wantSteps:     []string{"init", "step2"},
			wantErr:       false,
		},
		{
			name: "skip steps of same name in multiple stages",
			pipeline: &pipeline.Build{
				Stages: []*pipeline.Stage{
					{
						Name: "stage1",
						Steps: []*pipeline.Container{
							{Name: "init"},
							{Name: "step1"},
							{Name: "step2"},
						},
					},
					{
						Name: "stage2",
						Steps: []*pipeline.Container{
							{Name: "step1"},
							{Name: "step2"},
						},
					},
				},
			},
			stepsToRemove: []string{"step1"},
			wantSteps:     []string{"init", "step2", "step2"},
			wantErr:       false,
		},
		{
			name: "skip all steps in all stages except init",
			pipeline: &pipeline.Build{
				Stages: []*pipeline.Stage{
					{
						Name: "stage1",
						Steps: []*pipeline.Container{
							{Name: "init"},
						},
					},
					{
						Name: "stage2",
						Steps: []*pipeline.Container{
							{Name: "step1"},
						},
					},
					{
						Name: "stage3",
						Steps: []*pipeline.Container{
							{Name: "step1"},
						},
					},
				},
			},
			stepsToRemove: []string{"step1"},
			wantSteps:     []string{"init"},
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := skipSteps(tt.pipeline, tt.stepsToRemove)
			if (err != nil) != tt.wantErr {
				t.Errorf("skipSteps() error = %v, wantErr %v", err, tt.wantErr)
			}

			// ensure the pipeline has the right state after step removal
			var remainingSteps []string

			if len(tt.pipeline.Steps) > 0 {
				for _, step := range tt.pipeline.Steps {
					remainingSteps = append(remainingSteps, step.Name)
				}
			} else if len(tt.pipeline.Stages) > 0 {
				for _, stage := range tt.pipeline.Stages {
					for _, step := range stage.Steps {
						remainingSteps = append(remainingSteps, step.Name)
					}
				}
			}

			if !reflect.DeepEqual(remainingSteps, tt.wantSteps) {
				t.Errorf("remaining steps = %v, want %v", remainingSteps, tt.wantSteps)
			}
		})
	}
}

func TestFormatStepIdentifier(t *testing.T) {
	tests := []struct {
		name      string
		stageName string
		stepName  string
		want      string
	}{
		{
			name:      "basic format",
			stageName: "build",
			stepName:  "test",
			want:      "[stage: build][step: test]",
		},
		{
			name:      "empty stage name",
			stageName: "",
			stepName:  "test",
			want:      "[step: test]",
		},
		{
			name:      "empty stage and step name",
			stageName: "",
			stepName:  "",
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatStepIdentifier(tt.stageName, tt.stepName); got != tt.want {
				t.Errorf("formatStepIdentifier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func init() {
	// discard logs for tests
	logrus.SetOutput(io.Discard)
}
