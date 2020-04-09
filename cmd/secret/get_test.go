// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

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

var testSecretAppGet = cli.NewApp()

// setup the command for tests
func init() {
	testSecretAppGet.Commands = []*cli.Command{
		{
			Name: "get",
			Subcommands: []*cli.Command{
				&GetCmd,
			},
		},
	}
	testSecretAppGet.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "addr",
		},
		&cli.StringFlag{
			Name: "token",
		},
	}
}

func TestSecret_Get_Success(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testSecretAppGet, set, nil)

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
			"get", "secret",
			// "--engine", "native", "--type", "repository",
			"--org", "github", "--repo", "octocat"}, want: nil},

		// json output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "secret",
			"--engine", "native", "--type", "repository",
			"--org", "github", "--repo", "octocat", "--o", "json"}, want: nil},

		// yaml output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "secret",
			"--engine", "native", "--type", "repository",
			"--org", "github", "--repo", "octocat", "--o", "yaml"}, want: nil},

		// wide output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "secret",
			"--engine", "native", "--type", "repository",
			"--org", "github", "--repo", "octocat", "--o", "wide"}, want: nil},

		// page default output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "secret",
			"--engine", "native", "--type", "repository",
			"--org", "github", "--repo", "octocat", "--p", "2"}, want: nil},

		// per page default output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "secret",
			"--engine", "native", "--type", "repository",
			"--org", "github", "--repo", "octocat", "--pp", "20"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testSecretAppGet.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestSecret_Get_Failure(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testSecretAppGet, set, nil)

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
			"get", "secret",
			"--engine", "native", "--type", "repository",
			"--org", "github", "--repo", "octocat"},
			want: fmt.Errorf("Invalid command: Flag '--addr' is not set or is empty")},

		// ´Error with invalid token
		{data: []string{
			"", "--addr", s.URL,
			"get", "secret",
			"--engine", "native", "--type", "repository",
			"--org", "github", "--repo", "octocat"},
			want: fmt.Errorf("Invalid command: Flag '--token' is not set or is empty")},

		// ´Error with invalid org
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "secret",
			"--engine", "native", "--type", "repository",
			"--repo", "octocat"},
			want: fmt.Errorf("Invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid repo
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "secret",
			"--engine", "native", "--type", "repository",
			"--org", "github"},
			want: fmt.Errorf("Invalid command: Flag '--repo' is not set or is empty")},
	}

	// run test
	for _, test := range tests {
		got := testSecretAppGet.Run(test.data)

		if got == test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
