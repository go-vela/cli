// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"flag"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

var testDeploymentAppAdd = cli.NewApp()

// setup the command for tests
func init() {
	testDeploymentAppAdd.Commands = []*cli.Command{
		{
			Name: "add",
			Subcommands: []*cli.Command{
				&AddCmd,
			},
		},
	}
	testDeploymentAppAdd.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "addr",
		},
		&cli.StringFlag{
			Name: "token",
		},
	}
}

func TestDeployment_Add_Success(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testDeploymentAppAdd, set, nil)

	// setup server
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())

	// setup types
	tests := []struct {
		data []string
		want error
	}{
		{ // Add a deployment for a repo.
			data: []string{
				"", "--addr", s.URL,
				"--token", "foobar",
				"add",
				"deployment",
				"--org", "github",
				"--repo", "octocat",
			},
			want: nil,
		},
		{ // Add a deployment with specific target environment.
			data: []string{"",
				"--addr", s.URL,
				"--token", "foobar",
				"add",
				"deployment",
				"--org", "github",
				"--repo", "octocat",
				"--target", "stage",
			},
			want: nil,
		},
		{ // Add a deployment with a specific reference.
			data: []string{"",
				"--addr", s.URL,
				"--token", "foobar",
				"add",
				"deployment",
				"--org", "github",
				"--repo", "octocat",
				"--ref", "dev",
			},
			want: nil,
		},
	}

	// run test
	for _, test := range tests {
		got := testDeploymentAppAdd.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestDeployment_Add_Failure(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testDeploymentAppAdd, set, nil)

	// setup server
	gin.SetMode(gin.TestMode)

	s := httptest.NewServer(server.FakeHandler())

	// setup types
	tests := []struct {
		data []string
		want error
	}{
		{ // Error with invalid org
			data: []string{"",
				"--token", "foobar",
				"add",
				"deployment",
				"--repo", "octocat",
			},
			want: fmt.Errorf("Invalid command: Flag '--org' is not set or is empty"),
		},
		{ // Error with invalid repo
			data: []string{"",
				"--addr", s.URL,
				"--token", "foobar",
				"add",
				"deployment",
				"--org", "github",
			},
			want: fmt.Errorf("Invalid command: Flag '--repo' is not set or is empty"),
		},
		{ // Error with invalid ref
			data: []string{"",
				"--addr", s.URL,
				"--token", "foobar",
				"add",
				"deployment",
				"--org", "github",
				"--repo", "octocat",
				"--ref", "",
			},
			want: fmt.Errorf("Invalid command: Flag '--ref' is not set or is empty"),
		},
		{ // Error with invalid target
			data: []string{"",
				"--addr", s.URL,
				"--token", "foobar",
				"add",
				"deployment",
				"--org", "github",
				"--repo", "octocat",
				"--target", "",
			},
			want: fmt.Errorf("Invalid command: Flag '--target' is not set or is empty"),
		},
	}

	// run test
	for _, test := range tests {
		got := testDeploymentAppAdd.Run(test.data)

		if got == test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
