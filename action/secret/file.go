// SPDX-License-Identifier: Apache-2.0

package secret

import "github.com/go-vela/types/library"

// ConfigFile represents the secret configuration necessary
// to perform secret related requests from a file with Vela.
type ConfigFile struct {
	Metadata struct {
		Version string `yaml:"version,omitempty"`
		Engine  string `yaml:"engine,omitempty"`
	} `yaml:"metadata,omitempty"`
	Secrets []*library.Secret `yaml:"secrets,omitempty"`
}
