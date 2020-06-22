// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package hook

import (
	"sort"
	"time"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/types/library"

	"github.com/dustin/go-humanize"
	"github.com/gosuri/uitable"
)

// table is a helper function to output the
// provided hooks in a table format with
// a specific set of fields displayed.
func table(hooks *[]library.Hook) error {
	// create new table
	table := uitable.New()

	// set column width for table to 50
	table.MaxColWidth = 50

	// ensure the table is always wrapped
	table.Wrap = true

	// set of hook fields we display in a table
	table.AddRow("NUMBER", "STATUS", "EVENT", "BRANCH", "CREATED")

	// iterate through all hooks in the list
	for _, h := range reverse(*hooks) {
		// calculate created timestamp in human readable form
		c := humanize.Time(time.Unix(h.GetCreated(), 0))

		// add a row to the table with the specified values
		table.AddRow(h.GetNumber(), h.GetStatus(), h.GetEvent(), h.GetBranch(), c)
	}

	// output the table in default format
	err := output.Default(table)
	if err != nil {
		return err
	}

	return nil
}

// wideTable is a helper function to output the
// provided hooks in a wide table format with
// a specific set of fields displayed.
func wideTable(hooks *[]library.Hook) error {
	// create new wide table
	table := uitable.New()

	// set column width for wide table to 200
	table.MaxColWidth = 200

	// ensure the wide table is always wrapped
	table.Wrap = true

	// set of hook fields we display in a table
	table.AddRow("NUMBER", "SOURCE", "STATUS", "HOST", "EVENT", "BRANCH", "CREATED")

	// iterate through all hooks in the list
	for _, h := range reverse(*hooks) {
		// calculate created timestamp in human readable form
		c := humanize.Time(time.Unix(h.GetCreated(), 0))

		// add a row to the table with the specified values
		table.AddRow(h.GetNumber(), h.GetSourceID(), h.GetStatus(), h.GetHost(), h.GetEvent(), h.GetBranch(), c)
	}

	// output the wide table in default format
	err := output.Default(table)
	if err != nil {
		return err
	}

	return nil
}

// reverse is a helper function to sort the hooks
// based off the hook number and then flip the
// order they get displayed in.
func reverse(h []library.Hook) []library.Hook {
	sort.SliceStable(h, func(i, j int) bool {
		return h[i].GetNumber() < h[j].GetNumber()
	})

	return h
}
