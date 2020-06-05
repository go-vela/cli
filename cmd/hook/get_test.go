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

var testHookAppGet = cli.NewApp()

// setup the command for tests
func init() {
	testHookAppGet.Commands = []*cli.Command{
		{
			Name: "get",
			Subcommands: []*cli.Command{
				&GetCmd,
			},
		},
	}
	testHookAppGet.Flags = []cli.Flag{
		&cli.StringFlag{
			Name: "addr",
		},
		&cli.StringFlag{
			Name: "token",
		},
	}
}

func TestHook_Get_Success(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testHookAppGet, set, nil)

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
			"get", "hooks",
			"--org", "github", "--repo", "octocat"}, want: nil},

		// json output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "hooks",
			"--org", "github", "--repo", "octocat", "--o", "json"}, want: nil},

		// yaml output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "hooks",
			"--org", "github", "--repo", "octocat", "--o", "yaml"}, want: nil},

		// wide output
		{data: []string{
			"", "--addr", s.URL, "--token", "foobar",
			"get", "hooks",
			"--org", "github", "--repo", "octocat", "--o", "wide"}, want: nil},
	}

	// run test
	for _, test := range tests {
		got := testHookAppGet.Run(test.data)

		if got != test.want {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}

func TestHook_Get_Failure(t *testing.T) {

	set := flag.NewFlagSet("test", 0)
	_ = cli.NewContext(testHookAppGet, set, nil)

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
			"get", "hooks",
			"--repo", "octocat"},
			want: fmt.Errorf("invalid command: Flag '--org' is not set or is empty")},

		// ´Error with invalid repo
		{data: []string{
			"", "--addr", s.URL,
			"get", "hooks",
			"--org", "github"},
			want: fmt.Errorf("invalid command: Flag '--repo' is not set or is empty")},
	}

	// run test
	for _, test := range tests {
		got := testHookAppGet.Run(test.data)
		if got == nil || got.Error() != test.want.Error() {
			t.Errorf("Run is %v, want %v", got, test.want)
		}
	}
}
