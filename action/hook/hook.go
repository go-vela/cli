// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package hook

// Config represents the configuration necessary
// to perform hook related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Number  int
	Page    int
	PerPage int
	Output  string
}
