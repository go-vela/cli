// SPDX-License-Identifier: Apache-2.0

package hook

import (
	"sort"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gosuri/uitable"
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/types/library"
)

// table is a helper function to output the
// provided hooks in a table format with
// a specific set of fields displayed.
func table(hooks *[]library.Hook) error {
	logrus.Debug("creating table for list of hooks")

	// create a new table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#New
	table := uitable.New()

	// set column width for table to 50
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table
	table.MaxColWidth = 50

	// ensure the table is always wrapped
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table
	table.Wrap = true

	logrus.Trace("adding headers to hook table")

	// set of hook fields we display in a table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("NUMBER", "STATUS", "EVENT", "BRANCH", "CREATED")

	// iterate through all hooks in the list
	for _, h := range reverse(*hooks) {
		logrus.Tracef("adding hook %d to hook table", h.GetNumber())

		// calculate created timestamp in human readable form
		//
		// https://pkg.go.dev/github.com/dustin/go-humanize?tab=doc#Time
		c := humanize.Time(time.Unix(h.GetCreated(), 0))

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(h.GetNumber(), h.GetStatus(), h.GetEvent(), h.GetBranch(), c)
	}

	// output the table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// wideTable is a helper function to output the
// provided hooks in a wide table format with
// a specific set of fields displayed.
func wideTable(hooks *[]library.Hook) error {
	logrus.Debug("creating wide table for list of hooks")

	// create new wide table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#New
	table := uitable.New()

	// set column width for wide table to 200
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table
	table.MaxColWidth = 200

	// ensure the wide table is always wrapped
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table
	table.Wrap = true

	logrus.Trace("adding headers to wide hook table")

	// set of hook fields we display in a wide table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("NUMBER", "SOURCE", "STATUS", "HOST", "EVENT", "BRANCH", "CREATED")

	// iterate through all hooks in the list
	for _, h := range reverse(*hooks) {
		logrus.Tracef("adding hook %d to wide hook table", h.GetNumber())

		// calculate created timestamp in human readable form
		//
		// https://pkg.go.dev/github.com/dustin/go-humanize?tab=doc#Time
		c := humanize.Time(time.Unix(h.GetCreated(), 0))

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(h.GetNumber(), h.GetSourceID(), h.GetStatus(), h.GetHost(), h.GetEvent(), h.GetBranch(), c)
	}

	// output the wide table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// reverse is a helper function to sort the hooks
// based off the hook number and then flip the
// order they get displayed in.
func reverse(h []library.Hook) []library.Hook {
	// sort the list of hooks based off the hook number
	sort.SliceStable(h, func(i, j int) bool {
		return h[i].GetNumber() < h[j].GetNumber()
	})

	return h
}
