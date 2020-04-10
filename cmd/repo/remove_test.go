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

var testRepoAppRemove = cli.NewApp()

// setup the command for tests
func init() {
	testRepoAppRemove.Commands = []*cli.Command{
		{
			Name: "remove",
			Subcommands: []*cli.Command{
				&RemoveCmd,
			},
		},
	}
	testRepoAppRemove.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "addr",
		},
		&cli.StringFlag{
			Name: "token",
		},
	}
}

func TestRepo_Remove_Success(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testRepoAppRemove, set, nil)

	// setup server
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())

	// setup types
	tests := []struct {
		data []string
		want error
	}{
		// remove a repository
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"remove", "repo", "--org", "github", "--repo", "octocat"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testRepoAppRemove.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestRepo_Remove_Failure(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testRepoAppRemove, set, nil)

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
			"remove", "repo", "--repo", "octocat",
			"--event", "push"},
			want: fmt.Errorf("Invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid repo
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"remove", "repo", "--org", "github",
			"--event", "push"},
			want: fmt.Errorf("Invalid command: Flag '--repo' is not set or is empty")},
	}

	// run test
	for _, test := range tests {
		got := testRepoAppRemove.Run(test.data)

		if got == test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
