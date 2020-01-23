// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"flag"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

var testRepoAppAdd = cli.NewApp()

// setup the command for tests
func init() {
	testRepoAppAdd.Commands = []cli.Command{
		{
			Name: "add",
			Subcommands: []cli.Command{
				AddCmd,
			},
		},
	}
	testRepoAppAdd.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "addr",
		},
		cli.StringFlag{
			Name: "token",
		},
	}
}

func TestRepo_Add_Success(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testRepoAppAdd, set, nil)

	// setup server
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())

	// setup types
	tests := []struct {
		data []string
		want error
	}{
		// Add a repository with push and pull request enabled
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"add", "repo", "--org", "github", "--repo", "octocat",
			"--event", "push", "--event", "pull_request"}, want: nil},

		// Add a repository with all event types enabled
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"add", "repo", "--org", "github", "--repo", "octocat",
			"--event", "push", "--event", "pull_request",
			"--event", "tag", "--event", "deployment"}, want: nil},

		// Add a repository with a longer build timeout
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"add", "repo", "--org", "github", "--repo", "octocat",
			"--event", "push", "--timeout", "90"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testRepoAppAdd.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestRepo_Add_Failure(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testRepoAppAdd, set, nil)

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
			"", "--token", "foobar",
			"add", "repo", "--repo", "octocat",
			"--event", "push"},
			want: fmt.Errorf("Invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid repo
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"add", "repo", "--org", "github",
			"--event", "push"},
			want: fmt.Errorf("Invalid command: Flag '--repo' is not set or is empty")},
	}

	// run test
	for _, test := range tests {
		got := testRepoAppAdd.Run(test.data)

		if got == test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
