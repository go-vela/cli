// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

import (
	"flag"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var testLogAppView = cli.NewApp()

// setup the command for tests
func init() {
	testLogAppView.Commands = []*cli.Command{
		{
			Name: "view",
			Subcommands: []*cli.Command{
				&ViewCmd,
			},
		},
	}
	testLogAppView.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "addr",
		},
		&cli.StringFlag{
			Name: "token",
		},
	}
}

func TestLog_View_Success(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testLogAppView, set, nil)

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
			"view", "log",
			"--org", "github", "--repo", "octocat", "--b", "1", "--t", "step"}, want: nil},
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "log",
			"--org", "github", "--repo", "octocat", "--b", "1", "--t", "service"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testLogAppView.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestLog_View_Failure(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testLogAppView, set, nil)

	// setup server
	gin.SetMode(gin.TestMode)
	s := httptest.NewServer(server.FakeHandler())

	// setup types
	tests := []struct {
		data []string
		want error
	}{

		// ´Error with invalid org
		{data: []string{
			"", "--addr", s.URL,
			"view", "log",
			"--repo", "octocat", "--b", "1"},
			want: fmt.Errorf("invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid repo
		{data: []string{
			"", "--addr", s.URL,
			"view", "log",
			"--org", "github", "--b", "1"},
			want: fmt.Errorf("invalid command: Flag '--repo' is not set or is empty")},

		// ´Error with invalid number
		{data: []string{
			"", "--addr", s.URL,
			"view", "log",
			"--org", "github", "--repo", "octocat"},
			want: fmt.Errorf("invalid command: Flag '--build-number' is not set or is empty")},

		// ´Error with type
		{data: []string{
			"", "--addr", s.URL,
			"view", "log",
			"--org", "github", "--repo", "octocat", "--b", "1"},
			want: fmt.Errorf("invalid command: Flag '--type' is not set or is empty")},

		// ´Error with type
		{data: []string{
			"", "--addr", s.URL,
			"view", "log",
			"--org", "github", "--repo", "octocat", "--b", "1", "--t", "invalid"},
			want: fmt.Errorf("invalid command: Flag '--type' is not valid")},
	}

	// run test
	for _, test := range tests {
		got := testLogAppView.Run(test.data)
		if got == nil || got.Error() != test.want.Error() {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
