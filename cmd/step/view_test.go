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

var testStepAppView = cli.NewApp()

// setup the command for tests
func init() {
	testStepAppView.Commands = []cli.Command{
		{
			Name: "view",
			Subcommands: []cli.Command{
				ViewCmd,
			},
		},
	}
	testStepAppView.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "addr",
		},
		cli.StringFlag{
			Name: "token",
		},
	}
}

func TestStep_View_Success(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testStepAppView, set, nil)

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
			"view", "step",
			"--org", "github", "--repo", "octocat", "--build-number", "1", "--step-number", "1"}, want: nil},

		// json output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "step",
			"--org", "github", "--repo", "octocat", "--build-number", "1", "--step-number", "1", "--o", "json"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testStepAppView.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestStep_View_Failure(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testStepAppView, set, nil)

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
			"view", "step",
			"--org", "github", "--repo", "octocat", "--build-number", "1", "--step-number", "1"},
			want: fmt.Errorf("Invalid command: Flag '--addr' is not set or is empty")},

		// ´Error with invalid token
		{data: []string{
			"", "--addr", s.URL,
			"view", "step",
			"--org", "github", "--repo", "octocat", "--build-number", "1", "--step-number", "1"},
			want: fmt.Errorf("Invalid command: Flag '--token' is not set or is empty")},

		// ´Error with invalid org
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "step",
			"--repo", "octocat", "--build-number", "1", "--step-number", "1"},
			want: fmt.Errorf("Invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid repo
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "step",
			"--org", "github", "--build-number", "1", "--step-number", "1"},
			want: fmt.Errorf("Invalid command: Flag '--repo' is not set or is empty")},

		// ´Error with invalid build-number
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "step",
			"--org", "github", "--repo", "octocat", "--step-number", "1"},
			want: fmt.Errorf("Invalid command: Flag '--build-number' is not set or is empty")},

		// ´Error with invalid step-number
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "step",
			"--org", "github", "--repo", "octocat", "--build-number", "1"},
			want: fmt.Errorf("Invalid command: Flag '--step-number' is not set or is empty")},
	}

	// run test
	for _, test := range tests {
		got := testStepAppView.Run(test.data)

		if got == test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
