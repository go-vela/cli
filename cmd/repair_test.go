// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package cmd

import (
	"flag"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var testAppRepair = cli.NewApp()

// setup the command for tests
func init() {
	testAppRepair.Commands = []*cli.Command{
		&repairCmd,
	}

	testAppRepair.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "addr",
		},
		&cli.StringFlag{
			Name: "token",
		},
	}
}

func TestRepo_Repair_Success(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testAppRepair, set, nil)

	// setup server
	gin.SetMode(gin.TestMode)
	s := httptest.NewServer(server.FakeHandler())

	// setup types
	tests := []struct {
		data []string
		want error
	}{
		// repair a repository
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"repair", "--org", "github", "--repo", "octocat"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testAppRepair.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestRepo_Repair_Failure(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testAppRepair, set, nil)

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
			"repair", "--repo", "octocat"},
			want: fmt.Errorf("Invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid repo
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"repair", "--org", "github"},
			want: fmt.Errorf("Invalid command: Flag '--repo' is not set or is empty")},
	}

	// run test
	for _, test := range tests {
		got := testAppRepair.Run(test.data)

		if got == test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
