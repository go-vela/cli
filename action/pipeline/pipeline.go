// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

// Config represents the configuration necessary
// to perform pipeline related requests with Vela.
type Config struct {
	Action   string
	File     string
	Path     string
	Type     string
	Stages   bool
	Template bool
	Output   string
}
