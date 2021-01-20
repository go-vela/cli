// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

// Config represents the configuration necessary
// to perform build related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Number  int
	Page    int
	PerPage int
	Output  string
}
