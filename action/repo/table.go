// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package repo

import (
	"strings"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"

	"github.com/gosuri/uitable"
)

// table is a helper function to output the
// provided repos in a table format with
// a specific set of fields displayed.
func table(repos *[]library.Repo) error {
	// create new table
	table := uitable.New()

	// set column width for table to 50
	table.MaxColWidth = 50

	// ensure the table is always wrapped
	table.Wrap = true

	// set of repo fields we display in a table
	table.AddRow("ORG/REPO", "ACTIVE", "EVENTS", "VISIBILITY", "BRANCH")

	// iterate through all repos in the list
	for _, r := range *repos {
		e := strings.Join(events(&r), ",")

		// add a row to the table with the specified values
		table.AddRow(r.GetFullName(), r.GetActive(), e, r.GetVisibility(), r.GetBranch())
	}

	// output the table in default format
	err := output.Default(table)
	if err != nil {
		return err
	}

	return nil
}

// wideTable is a helper function to output the
// provided repos in a wide table format with
// a specific set of fields displayed.
func wideTable(repos *[]library.Repo) error {
	// create new wide table
	table := uitable.New()

	// set column width for wide table to 200
	table.MaxColWidth = 200

	// ensure the wide table is always wrapped
	table.Wrap = true

	// set of repo fields we display in a wide table
	table.AddRow("ORG/REPO", "ACTIVE", "EVENTS", "VISIBILITY", "BRANCH", "REMOTE")

	// iterate through all repos in the list
	for _, r := range *repos {
		e := strings.Join(events(&r), ",")

		// add a row to the table with the specified values
		table.AddRow(r.GetFullName(), r.GetActive(), e, r.GetVisibility(), r.GetBranch(), r.GetLink())
	}

	// output the wide table in default format
	err := output.Default(table)
	if err != nil {
		return err
	}

	return nil
}

// events is a helper function to output
// the event types a repo is configured
// to trigger builds off of.
func events(r *library.Repo) []string {
	e := []string{}

	if r.GetAllowComment() {
		e = append(e, constants.EventComment)
	}

	if r.GetAllowDeploy() {
		e = append(e, constants.EventDeploy)
	}

	if r.GetAllowPull() {
		e = append(e, constants.EventPull)
	}

	if r.GetAllowPush() {
		e = append(e, constants.EventPush)
	}

	if r.GetAllowTag() {
		e = append(e, constants.EventTag)
	}

	return e
}
