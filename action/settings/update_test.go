// SPDX-License-Identifier: Apache-2.0

package settings

import (
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	"github.com/go-vela/server/mock/server"
)

func TestSettings_Config_Update(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}

	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action: "update",
				Output: "",

				Queue: Queue{
					Routes: &[]string{"test"},
				},
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Output: "",
				Compiler: Compiler{
					CloneImage:        vela.String("test"),
					TemplateDepth:     vela.Int(1),
					StarlarkExecLimit: vela.UInt64(1),
				},
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Update(client)

		if test.failure {
			if err == nil {
				t.Errorf("Update should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("Update returned err: %v", err)
		}
	}
}

func TestSettings_Config_UpdateFromFile(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}

	// setup tests
	tests := []struct {
		failure bool
		config  *Config
	}{
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/all.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/platform.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/compiler.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/queue.yml",
				Output: "",
			},
		},
		{
			failure: true,
			config: &Config{
				Action: "update",
				File:   "testdata/noplatform.yml",
				Output: "",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.UpdateFromFile(client)

		if test.failure {
			if err == nil {
				t.Errorf("UpdateFromFile should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("UpdateFromFile returned err: %v", err)
		}
	}
}
