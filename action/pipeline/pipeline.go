// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

// Config represents the configuration necessary
// to perform pipeline related requests with Vela.
type Config struct {
	Action      string
	Branch      string
	Comment     string
	Event       string
	Tag         string
	Target      string
	Org         string
	Repo        string
	Ref         string
	RawPipeline []byte
	File        string
	Path        string
	Type        string
	Stages      bool
	Template    bool
	Local       bool
	Volumes     []string
	Output      string
}
