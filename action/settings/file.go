// SPDX-License-Identifier: Apache-2.0

package settings

import "github.com/go-vela/server/api/types/settings"

// ConfigFile represents the configuration necessary
// to perform settings related requests from a file with Vela.
type ConfigFile struct {
	*settings.Platform `yaml:"platform,omitempty"`
}
