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

var testSecretAppUpdate = cli.NewApp()

// setup the command for tests
func init() {
	testSecretAppUpdate.Commands = []cli.Command{
		{
			Name: "update",
			Subcommands: []cli.Command{
				UpdateCmd,
			},
		},
	}
	testSecretAppUpdate.Flags = []cli.Flag{
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

func TestSecret_Update_Success(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testSecretAppUpdate, set, nil)

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
			"update", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--name", "foo"}, want: nil},

		// default repo output with global config
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar", "--engine", "native", "--type", "repo",
			"update", "secret",
			"--org", "github", "--repo", "octocat", "--name", "foo"}, want: nil},

		// default org output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"update", "secret",
			"--engine", "native", "--type", "org",
			"--org", "github", "--repo", "*", "--name", "foo"}, want: nil},

		// default org output with global config
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar", "--engine", "native", "--type", "org",
			"update", "secret",
			"--org", "github", "--repo", "*", "--name", "foo"}, want: nil},

		// default shared output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"update", "secret",
			"--engine", "native", "--type", "shared",
			"--org", "github", "--team", "octokitties", "--name", "foo"}, want: nil},

		// default shared output with global config
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar", "--engine", "native",
			"update", "secret",
			"--type", "shared", "--org", "github", "--team", "octokitties", "--name", "foo"}, want: nil},

		//TODO: Add test for file workflow
	}

	// run test
	for _, test := range tests {
		got := testSecretAppUpdate.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestSecret_Update_Failure(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testSecretAppUpdate, set, nil)

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
			"update", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--name", "foo"},
			want: fmt.Errorf("Invalid command: Flag '--addr' is not set or is empty")},

		// ´Error with invalid token
		{data: []string{
			"", "--addr", s.URL,
			"update", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat", "--name", "foo"},
			want: fmt.Errorf("Invalid command: Flag '--token' is not set or is empty")},

		// ´Error with invalid org
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"update", "secret",
			"--engine", "native", "--type", "repo",
			"--repo", "octocat", "--name", "foo"},
			want: fmt.Errorf("Invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid name
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"update", "secret",
			"--engine", "native", "--type", "repo",
			"--org", "github", "--repo", "octocat"},
			want: fmt.Errorf("Invalid command: Flag '--name' is not set or is empty")},
	}

	// run test
	for _, test := range tests {
		got := testSecretAppUpdate.Run(test.data)

		if got == test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
