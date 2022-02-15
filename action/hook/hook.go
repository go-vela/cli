// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package hook

// Config represents the configuration necessary
// to perform hook related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Output  string
	Number  int
	Page    int
	PerPage int
}
