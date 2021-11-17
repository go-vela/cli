// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package internal

import (
	"strings"

	"github.com/go-git/go-git/v5"
	giturls "github.com/whilp/git-urls"
)

func GetGitConfigOrg(path string) string {
	r, err := git.PlainOpen(path)
	if err != nil {
		return ""
	}
	origin, err := r.Remote("origin")
	if err != nil {
		return ""
	}
	url, err := giturls.Parse(origin.Config().URLs[0])
	if err != nil {
		return ""
	}
	org := strings.Split(url.Path, "/")[0]

	return org
}

func GetGitConfigRepo(path string) string {
	r, err := git.PlainOpen(path)
	if err != nil {
		return ""
	}
	origin, err := r.Remote("origin")
	if err != nil {
		return ""
	}
	url, err := giturls.Parse(origin.Config().URLs[0])
	if err != nil {
		return ""
	}
	repoName := strings.Split(strings.Split(url.Path, "/")[1], ".git")[0]
	return repoName
}
