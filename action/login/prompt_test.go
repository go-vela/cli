// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// +build !race

package login

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestLogin_Config_PromptUsername(t *testing.T) {
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
			data: "foo\n",
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
			data: "foo",
		},
	}

	// run tests
	for _, test := range tests {
		in, err := ioutil.TempFile("/tmp", "username")
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

		err = test.config.PromptUsername(in)

		if test.failure {
			if err == nil {
				t.Errorf("PromptUsername should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("PromptUsername returned err: %v", err)
		}
	}
}

func TestLogin_Config_PromptPassword(t *testing.T) {
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
			data: "foo\n",
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
			data: "foo",
		},
	}

	// run tests
	for _, test := range tests {
		in, err := ioutil.TempFile("/tmp", "password")
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

		err = test.config.PromptPassword(in)

		if test.failure {
			if err == nil {
				t.Errorf("PromptPassword should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("PromptPassword returned err: %v", err)
		}
	}
}

func TestLogin_Config_PromptOTP(t *testing.T) {
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
			data: "foo\n",
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
			data: "foo",
		},
	}

	// run tests
	for _, test := range tests {
		in, err := ioutil.TempFile("/tmp", "otp")
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

		err = test.config.PromptOTP(in)

		if test.failure {
			if err == nil {
				t.Errorf("PromptOTP should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("PromptOTP returned err: %v", err)
		}
	}
}
