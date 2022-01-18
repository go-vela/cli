// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package completion

// Config represents the configuration necessary
// to perform completion related requests with Vela.
type Config struct {
	Action string
	Bash   bool
	Zsh    bool
}
