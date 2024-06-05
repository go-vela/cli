// SPDX-License-Identifier: Apache-2.0

package deployment

import (
	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/types/raw"
)

// Config represents the configuration necessary
// to perform deployment related quests with Vela.
type Config struct {
	Action      string
	Org         string
	Repo        string
	Number      int
	Description string
	Ref         string
	Target      string
	Task        string
	Page        int
	PerPage     int
	Output      string
	Color       output.ColorOptions
	Parameters  raw.StringSliceMap
}
