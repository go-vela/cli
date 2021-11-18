// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package config

// ConfigFile represents the configuration file
// to perform config related requests with Vela.
//
// nolint: golint // ignore studder for package and struct name
type ConfigFile struct {
	API      *API      `yaml:"api,omitempty"`
	Log      *Log      `yaml:"log,omitempty"`
	GitSync  string    `yaml:"git_sync,omitempty"`
	Secret   *Secret   `yaml:"secret,omitempty"`
	Compiler *Compiler `yaml:"compiler,omitempty"`
	Output   string    `yaml:"output,omitempty"`
	Org      string    `yaml:"org,omitempty"`
	Repo     string    `yaml:"repo,omitempty"`
}

// API represents the API related configuration fields
// populated in the config file to perform requests
// with Vela.
type API struct {
	Address      string `yaml:"addr,omitempty"`
	Token        string `yaml:"token,omitempty"`
	AccessToken  string `yaml:"access_token,omitempty"`
	RefreshToken string `yaml:"refresh_token,omitempty"`
	Version      string `yaml:"version,omitempty"`
}

// Log represents the log related configuration fields
// populated in the config file to perform requests
// with Vela.
type Log struct {
	Level string `yaml:"level,omitempty"`
}

// Secret represents the secret configuration fields
// populated in the config file to perform requests
// with Vela.
type Secret struct {
	Engine string `yaml:"engine,omitempty"`
	Type   string `yaml:"type,omitempty"`
}

// Compiler represents the compiler configuration fields
// populated in the config file to perform requests
// with Vela.
type Compiler struct {
	GitHub *GitHub `yaml:"github,omitempty"`
}

// GitHub represents the compiler configuration fields
// populated in the config file to perform requests
// with Vela.
type GitHub struct {
	Token string `yaml:"token,omitempty"`
	URL   string `yaml:"url,omitempty"`
}
