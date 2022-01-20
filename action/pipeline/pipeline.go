// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package pipeline

// Config represents the configuration necessary
// to perform pipeline related requests with Vela.
type Config struct {
	Action        string
	Branch        string
	Comment       string
	Event         string
	Tag           string
	Target        string
	Org           string
	Repo          string
	Ref           string
	File          string
	Path          string
	Type          string
	Stages        bool
	Template      bool
	TemplateFiles []string
	Local         bool
	Remote        bool
	Volumes       []string
	Output        string
	PipelineType  string
}
