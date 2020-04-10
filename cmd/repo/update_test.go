// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"flag"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-vela/mock/server"
	"github.com/urfave/cli/v2"
)

var testRepoAppUpdate = cli.NewApp()

// setup the command for tests
func init() {
	testRepoAppUpdate.Commands = []*cli.Command{
		{
			Name: "update",
			Subcommands: []*cli.Command{
				&UpdateCmd,
			},
		},
	}
	testRepoAppUpdate.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "addr",
		},
		&cli.StringFlag{
			Name: "token",
		},
	}
}

func TestRepo_Update_Success(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testRepoAppUpdate, set, nil)

	// setup server
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())

	// setup types
	tests := []struct {
		data []string
		want error
	}{
		// Update a repository
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"update", "repo", "--org", "github", "--repo", "octocat"}, want: nil},

		// Update a repository with all event types enabled
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"update", "repo", "--org", "github", "--repo", "octocat",
			"--event", "push", "--event", "pull_request",
			"--event", "tag", "--event", "deployment", "--event", "comment"}, want: nil},

		// Update a repository with a longer build timeout
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"update", "repo", "--org", "github", "--repo", "octocat",
			"--event", "push", "--timeout", "90"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testRepoAppUpdate.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestRepo_Update_Failure(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testRepoAppUpdate, set, nil)

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
			"update", "repo", "--repo", "octocat",
			"--event", "push"},
			want: fmt.Errorf("Invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid repo
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"update", "repo", "--org", "github",
			"--event", "push"},
			want: fmt.Errorf("Invalid command: Flag '--repo' is not set or is empty")},
	}

	// run test
	for _, test := range tests {
		got := testRepoAppUpdate.Run(test.data)

		if got == test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
