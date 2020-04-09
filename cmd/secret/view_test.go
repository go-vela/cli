// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this secret.

package secret

import (
	"flag"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/mock/server"
	"github.com/urfave/cli/v2"
)

var testSecretAppView = cli.NewApp()

// setup the command for tests
func init() {
	testSecretAppView.Commands = []*cli.Command{
		{
			Name: "view",
			Subcommands: []*cli.Command{
				&ViewCmd,
			},
		},
	}
	testSecretAppView.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "addr",
		},
		&cli.StringFlag{
			Name: "token",
		},
		&cli.StringFlag{
			Name: "engine",
		},
		&cli.StringFlag{
			Name: "type",
		},
	}
}

func TestSecret_View_Success(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testSecretAppView, set, nil)

	// setup server
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())

	// setup types
	tests := []struct {
		data []string
		want error
	}{
		// default repo output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--name", "foo"}, want: nil},

		// default repo output with global config
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar", "--engine", "native", "--type", "repo",
			"view", "secret",
			"--org", "github", "--repo", "octocat", "--name", "foo"}, want: nil},

		// repo json output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--name", "*", "--output", "json"}, want: nil},

		// default org output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "secret",
			"--engine", "native", "--type", "org",
			"--org", "github", "--repo", "*", "--name", "foo"}, want: nil},

		// default org output with global config
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar", "--engine", "native", "--type", "org",
			"view", "secret",
			"--org", "github", "--repo", "*", "--name", "foo"}, want: nil},

		// org json output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "secret",
			"--engine", "native", "--type", "org",
			"--org", "github", "--repo", "*", "--name", "foo", "--output", "json"}, want: nil},

		// default shared output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "secret",
			"--engine", "native", "--type", "shared",
			"--org", "github", "--team", "octokitties", "--name", "foo"}, want: nil},

		// default shared output with global config
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar", "--engine", "native", "--type", "shared",
			"view", "secret",
			"--org", "github", "--team", "octokitties", "--name", "foo"}, want: nil},

		// org shared output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "secret",
			"--engine", "native", "--type", "shared",
			"--org", "github", "--team", "octokitties", "--name", "foo", "--output", "json"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testSecretAppView.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestSecret_View_Failure(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testSecretAppView, set, nil)

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
			"view", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--name", "foo"},
			want: fmt.Errorf("Invalid command: Flag '--addr' is not set or is empty")},

		// ´Error with invalid token
		{data: []string{
			"", "--addr", s.URL,
			"view", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--name", "foo"},
			want: fmt.Errorf("Invalid command: Flag '--token' is not set or is empty")},

		// ´Error with invalid org
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "secret",
			"--engine", "native", "--type", "repo",
			"--repo", "octocat", "--name", "foo"},
			want: fmt.Errorf("Invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid name
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat"},
			want: fmt.Errorf("Invalid command: Flag '--name' is not set or is empty")},
	}

	// run test
	for _, test := range tests {
		got := testSecretAppView.Run(test.data)

		if got == test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
