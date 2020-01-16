// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package step

import (
	"flag"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/mock/server"
	"github.com/urfave/cli"
)

var testStepAppGet = cli.NewApp()

// setup the command for tests
func init() {
	testStepAppGet.Commands = []cli.Command{
		{
			Name: "get",
			Subcommands: []cli.Command{
				GetCmd,
			},
		},
	}
	testStepAppGet.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "addr",
		},
		cli.StringFlag{
			Name: "token",
		},
	}
}

func TestStep_Get_Success(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testStepAppGet, set, nil)

	// setup server
	gin.SetMode(gin.TestMode)
	s := httptest.NewServer(server.FakeHandler())

	// setup types
	tests := []struct {
		data []string
		want error
	}{
		// default output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "step",
			"--org", "github", "--repo", "octocat", "--build-number", "1"}, want: nil},

		// json output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "step",
			"--org", "github", "--repo", "octocat", "--build-number", "1", "--o", "json"}, want: nil},

		// yaml output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "step",
			"--org", "github", "--repo", "octocat", "--build-number", "1", "--o", "yaml"}, want: nil},

		// wide output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "step",
			"--org", "github", "--repo", "octocat", "--build-number", "1", "--o", "wide"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testStepAppGet.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestStep_Get_Failure(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testStepAppGet, set, nil)

	// setup server
	gin.SetMode(gin.TestMode)
	s := httptest.NewServer(server.FakeHandler())

	// setup types
	tests := []struct {
		data []string
		want error
	}{
		// ´Error with invalid addr
		{data: []string{
			"", "--token", "foobar",
			"get", "step",
			"--org", "github", "--repo", "octocat", "--build-number", "1"},
			want: fmt.Errorf("Invalid command: Flag '--addr' is not set or is empty")},

		// ´Error with invalid token
		{data: []string{
			"", "--addr", s.URL,
			"get", "step",
			"--org", "github", "--repo", "octocat", "--build-number", "1"},
			want: fmt.Errorf("Invalid command: Flag '--token' is not set or is empty")},

		// ´Error with invalid org
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "step",
			"--repo", "octocat", "--build-number", "1"},
			want: fmt.Errorf("Invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid org
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "step",
			"--org", "github", "--build-number", "1"},
			want: fmt.Errorf("Invalid command: Flag '--repo' is not set or is empty")},

		// ´Error with invalid build-number
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "step",
			"--org", "github", "--repo", "octocat"},
			want: fmt.Errorf("Invalid command: Flag '--build-number' is not set or is empty")},
	}

	// run test
	for _, test := range tests {
		got := testStepAppGet.Run(test.data)

		if got == test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
