// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package docs

// Config represents the configuration necessary
// to perform docs related requests with Vela.
type Config struct {
	Action   string
	Markdown bool
	Man      bool
}
