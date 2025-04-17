// SPDX-License-Identifier: Apache-2.0

package deployment

import (
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/urfave/cli/v3"

	"github.com/go-vela/cli/test"
	"github.com/go-vela/server/compiler/types/raw"
	"github.com/go-vela/server/mock/server"
)

func TestDeployment_Add(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup tests
	tests := []struct {
		failure bool
		cmd     *cli.Command
		args    []string
	}{
		{
			failure: false,
			cmd:     test.TestCommand(s.URL, add, CommandAdd.Flags),
			args:    []string{"--org", "Org-1", "--repo", "Repo-1", "--ref", "123abc", "--target", "dev"},
		},
		{
			failure: true,
			cmd:     test.TestCommand(s.URL, add, CommandAdd.Flags),
			args:    []string{"--org", "Org-1"},
		},
		{
			failure: true,
			cmd:     test.TestCommand(s.URL, add, nil),
		},
	}

	// run tests
	for _, test := range tests {
		err := test.cmd.Run(t.Context(), append([]string{"test"}, test.args...))

		if test.failure {
			if err == nil {
				t.Errorf("add should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("add returned err: %v", err)
		}
	}
}

func Test_parseKeyValue(t *testing.T) {
	type args struct {
		input []string
	}

	tests := []struct {
		name    string
		args    args
		want    raw.StringSliceMap
		wantErr bool
	}{
		{"valid input", args{input: []string{"foo=test1", "bar=test2"}}, raw.StringSliceMap{"foo": "test1", "bar": "test2"}, false},
		{"invalid input", args{input: []string{"foo=test1", "badinput"}}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseKeyValue(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseKeyValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseKeyValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseParamFile(t *testing.T) {
	tests := []struct {
		name    string
		file    string
		want    raw.StringSliceMap
		wantErr bool
	}{
		{
			name: "valid simple input JSON",
			file: "testdata/deploy-params-simple.json",
			want: raw.StringSliceMap{
				"one":   "two",
				"three": "four",
				"five":  "six",
			},
		},
		{
			name: "valid comma input JSON",
			file: "testdata/deploy-params-comma.json",
			want: raw.StringSliceMap{
				"greeting": "hello, there!",
				"farewell": "so long, partner!",
			},
		},
		{
			name: "valid input ENV",
			file: "testdata/deploy-params-env.env",
			want: raw.StringSliceMap{
				"USER": "VELA",
				"REPO": "CLI",
				"ORG":  "GO-VELA",
			},
		},
		{
			name:    "invalid input JSON bad type",
			file:    "testdata/deploy-params-bad-type.json",
			wantErr: true,
		},
		{
			name:    "invalid input JSON bad structure",
			file:    "testdata/deploy-params-bad-struct.json",
			wantErr: true,
		},
		{
			name:    "invalid input nonexistent file",
			file:    "testdata/does-not-exist.json",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseParamFile(tt.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseParamFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseParamFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
