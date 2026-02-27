// SPDX-License-Identifier: Apache-2.0

package secret

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/sdk-go/vela"
	api "github.com/go-vela/server/api/types"
	"github.com/go-vela/server/mock/server"
)

func TestSecret_Config_Update(t *testing.T) {
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
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Value:  "bar",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Engine: "native",
				Type:   "org",
				Org:    "github",
				Repo:   "*",
				Name:   "foo",
				Value:  "bar",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Engine: "native",
				Type:   "shared",
				Org:    "github",
				Team:   "octokitties",
				Name:   "foo",
				Value:  "bar",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Value:  "@testdata/foo.txt",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Value:  "bar",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Value:  "bar",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Value:  "bar",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				Engine: "native",
				Type:   "repo",
				Org:    "github",
				Repo:   "octocat",
				Name:   "foo",
				Value:  "bar",
				Output: "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.Update(t.Context(), client)

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

func TestSecret_Config_UpdateFromFile(t *testing.T) {
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
				File:   "testdata/repo.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/org.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/shared.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/multiple.yml",
				Output: "",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/repo.yml",
				Output: "dump",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/repo.yml",
				Output: "json",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/repo.yml",
				Output: "spew",
			},
		},
		{
			failure: false,
			config: &Config{
				Action: "update",
				File:   "testdata/repo.yml",
				Output: "yaml",
			},
		},
	}

	// run tests
	for _, test := range tests {
		err := test.config.UpdateFromFile(t.Context(), client)

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

func TestSecret_Config_Update_NoImageOrRepoAllowlist(t *testing.T) {
	// setup test server
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("expected PUT request, got %s", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("unable to read request body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		payload := map[string]json.RawMessage{}
		err = json.Unmarshal(body, &payload)
		if err != nil {
			t.Errorf("unable to decode request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if _, ok := payload["images"]; ok {
			t.Errorf("expected images to be omitted in update payload")
		}

		if _, ok := payload["repo_allowlist"]; ok {
			t.Errorf("expected repo_allowlist to be omitted in update payload")
		}

		secret := &api.Secret{}
		secret.SetOrg("github")
		secret.SetRepo("octocat")
		secret.SetName("foo")
		secret.SetType("repo")
		secret.SetValue("bar")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(secret)
	}))
	defer s.Close()

	// create a vela client
	client, err := vela.NewClient(s.URL, "vela", nil)
	if err != nil {
		t.Errorf("unable to create client: %v", err)
	}

	config := &Config{
		Action: "update",
		Engine: "native",
		Type:   "repo",
		Org:    "github",
		Repo:   "octocat",
		Name:   "foo",
		Value:  "bar",
		Output: "",
	}

	err = config.Update(t.Context(), client)
	if err != nil {
		t.Errorf("Update returned err: %v", err)
	}
}
