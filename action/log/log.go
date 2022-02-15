// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package log

// Config represents the configuration necessary
// to perform log related requests with Vela.
type Config struct {
	Action  string
	Org     string
	Repo    string
	Output  string
	Build   int
	Service int
	Step    int
}
