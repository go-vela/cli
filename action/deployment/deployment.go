// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

// Config represents the configuration necessary
// to perform deployment related quests with Vela.
type Config struct {
	Action      string
	Org         string
	Repo        string
	Number      int
	Description string
	Ref         string
	Target      string
	Task        string
	Page        int
	PerPage     int
	Output      string
}
