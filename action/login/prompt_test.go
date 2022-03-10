// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

//go:build !race
// +build !race

package login

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLogin_Config_PromptBrowserConfirm(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
		data    string
	}{
		{
			failure: false,
			config: &Config{
				Action: "login",
			},
			data: "y\n",
		},
		{
			failure: false,
			config: &Config{
				Action: "login",
			},
			data: "\n",
		},
		{
			failure: true,
			config: &Config{
				Action: "login",
			},
			data: "n\n",
		},
	}

	// run tests
	for _, test := range tests {
		in, err := ioutil.TempFile("/tmp", "browser")
		if err != nil {
			t.Errorf("unable to create temporary file: %v", err)
		}

		defer os.Remove(in.Name())

		_, err = in.Write([]byte(test.data))
		if err != nil {
			t.Errorf("unable to write content to temporary file: %v", err)
		}

		_, err = in.Seek(0, 0)
		if err != nil {
			t.Errorf("unable to seek temporary file: %v", err)
		}

		defer in.Close()

		err = test.config.PromptBrowserConfirm(in)

		if test.failure {
			if err == nil {
				t.Errorf("PromptBrowserConfirm should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("PromptBrowserConfirm returned err: %v", err)
		}
	}
}

func TestLogin_Config_PromptConfigConfirm(t *testing.T) {
	// setup tests
	tests := []struct {
		failure bool
		config  *Config
		data    string
	}{
		{
			failure: false,
			config: &Config{
				Action: "login",
			},
			data: "y\n",
		},
		{
			failure: true,
			config: &Config{
				Action: "login",
			},
			data: "\n",
		},
		{
			failure: true,
			config: &Config{
				Action: "login",
			},
			data: "n\n",
		},
	}

	// run tests
	for _, test := range tests {
		in, err := ioutil.TempFile("/tmp", "config")
		if err != nil {
			t.Errorf("unable to create temporary file: %v", err)
		}

		defer os.Remove(in.Name())

		_, err = in.Write([]byte(test.data))
		if err != nil {
			t.Errorf("unable to write content to temporary file: %v", err)
		}

		_, err = in.Seek(0, 0)
		if err != nil {
			t.Errorf("unable to seek temporary file: %v", err)
		}

		defer in.Close()

		err = test.config.PromptConfigConfirm(in)

		if test.failure {
			if err == nil {
				t.Errorf("PromptConfigConfirm should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("PromptConfigConfirm returned err: %v", err)
		}
	}
}
