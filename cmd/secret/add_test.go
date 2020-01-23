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
	"github.com/urfave/cli"
)

var testSecretAppAdd = cli.NewApp()

// setup the command for tests
func init() {
	testSecretAppAdd.Commands = []cli.Command{
		{
			Name: "create",
			Subcommands: []cli.Command{
				AddCmd,
			},
		},
	}
	testSecretAppAdd.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "addr",
		},
		cli.StringFlag{
			Name: "token",
		},
		cli.StringFlag{
			Name: "engine",
		},
		cli.StringFlag{
			Name: "type",
		},
	}
}

func TestSecret_Add_Success(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testSecretAppAdd, set, nil)

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
			"create", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--name", "foo", "--value", "bar"}, want: nil},

		// default repo output with global config
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar", "--engine", "native", "--type", "repo",
			"create", "secret",
			"--org", "github", "--repo", "octocat", "--name", "foo", "--value", "bar"}, want: nil},

		// default org output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"create", "secret",
			"--engine", "native", "--type", "org",
			"--org", "github", "--repo", "*", "--name", "foo", "--value", "bar"}, want: nil},

		// default org output with global config
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar", "--engine", "native", "--type", "org",
			"create", "secret",
			"--org", "github", "--repo", "*", "--name", "foo", "--value", "bar"}, want: nil},

		// default shared output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"create", "secret",
			"--engine", "native", "--type", "shared",
			"--org", "github", "--team", "octokitties", "--name", "foo", "--value", "bar"}, want: nil},

		// default shared output with global config
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar", "--engine", "native",
			"create", "secret",
			"--type", "shared", "--org", "github", "--team", "octokitties", "--name", "foo", "--value", "bar"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testSecretAppAdd.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestSecret_Add_Failure(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testSecretAppAdd, set, nil)

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
			"create", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--name", "foo", "--value", "bar"},
			want: fmt.Errorf("Invalid command: Flag '--addr' is not set or is empty")},

		// ´Error with invalid token
		{data: []string{
			"", "--addr", s.URL,
			"create", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--name", "foo"},
			want: fmt.Errorf("Invalid command: Flag '--token' is not set or is empty")},

		// ´Error with invalid org
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"create", "secret",
			"--engine", "native", "--type", "repo",
			"--repo", "octocat", "--name", "foo", "--value", "bar"},
			want: fmt.Errorf("Invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid name
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"create", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--value", "bar"},
			want: fmt.Errorf("Invalid command: Flag '--name' is not set or is empty")},

		// ´Error with invalid value
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"create", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--name", "foo"},
			want: fmt.Errorf("Invalid command: Flag '--name' is not set or is empty")},

		//TODO: Add test for file workflow
	}

	// run test
	for _, test := range tests {
		got := testSecretAppAdd.Run(test.data)

		if got == test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
