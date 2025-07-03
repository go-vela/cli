// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"strconv"
)

// StringToBool converts a string to a boolean.
// Needed for backwards compatibility with the CLI (urface 2.x vs 3.x)
// where it was possible to do c.Bool() on a flag that was a string.
func StringToBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}

	return b
}
