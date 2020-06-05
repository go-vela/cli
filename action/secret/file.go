// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import "github.com/go-vela/types/library"

// ConfigFile represents the secret configuration necessary
// to perform secret related quests from a file with Vela.
type ConfigFile struct {
	Metadata struct {
		Version string `yaml:"version,omitempty"`
		Engine  string `yaml:"engine,omitempty"`
	} `yaml:"metadata,omitempty"`
	Secrets []*library.Secret `yaml:"secrets,omitempty"`
}
