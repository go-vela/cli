// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

// Config represents the configuration necessary
// to perform deployment related quests with Vela.
type Config struct {
	Action      string
	Org         string
	Repo        string
	Description string
	Ref         string
	Target      string
	Task        string
	Output      string
	Parameters  []string
	Page        int
	PerPage     int
	Number      int
}
