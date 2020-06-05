package hook

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-vela/mock/server"
	"github.com/urfave/cli/v2"
	"net/http/httptest"
	"testing"
)

var testHookAppView = cli.NewApp()

// setup the command for tests
func init() {
	testHookAppView.Commands = []*cli.Command{
		{
			Name: "view",
			Subcommands: []*cli.Command{
				&ViewCmd,
			},
		},
	}
	testHookAppView.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "addr",
		},
		&cli.StringFlag{
			Name: "token",
		},
	}
}

func TestHook_View_Success(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testHookAppView, set, nil)

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
			"view", "hook",
			"--org", "github", "--repo", "octocat", "--hook-number", "1"}, want: nil},

		// json output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "hook",
			"--org", "github", "--repo", "octocat", "--hook-number", "1", "--o", "json"}, want: nil},

		// yaml output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "hook",
			"--org", "github", "--repo", "octocat", "--hook-number", "1", "--o", "yaml"}, want: nil},

		// wide output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"view", "hook",
			"--org", "github", "--repo", "octocat", "--hook-number", "1", "--o", "wide"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testHookAppView.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestHook_View_Failure(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testHookAppView, set, nil)

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
			"view", "hook",
			"--repo", "octocat", "--hook", "1"},
			want: fmt.Errorf("invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid repo
		{data: []string{
			"", "--addr", s.URL,
			"view", "hook",
			"--org", "github", "--hook", "1"},
			want: fmt.Errorf("invalid command: Flag '--repo' is not set or is empty")},

		// ´Error with invalid number
		{data: []string{
			"", "--addr", s.URL,
			"view", "hook",
			"--org", "github", "--repo", "octocat"},
			want: fmt.Errorf("invalid command: Flag '--hook-number' is not set or is empty")},
	}

	// run test
	for _, test := range tests {
		got := testHookAppView.Run(test.data)
		if got == nil || got.Error() != test.want.Error() {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
