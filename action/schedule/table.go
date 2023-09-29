// SPDX-License-Identifier: Apache-2.0

package schedule

import (
	"time"

	"github.com/dustin/go-humanize"
	"github.com/go-vela/cli/internal/output"
	"github.com/go-vela/types/library"
	"github.com/gosuri/uitable"
	"github.com/sirupsen/logrus"
)

// table is a helper function to output the
// provided schedules in a table format with
// a specific set of fields displayed.
func table(schedules *[]library.Schedule) error {
	logrus.Debug("creating table for list of schedules")

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

	logrus.Trace("adding headers to schedule table")

	// set of schedule fields we display in a table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("NAME", "ENTRY", "ACTIVE", "SCHEDULED_AT", "BRANCH")

	// iterate through all schedules in the list
	for _, s := range *schedules {
		logrus.Tracef("adding schedule %s to schedule table", s.GetName())

		// calculate scheduled_at timestamp in human readable form
		//
		// https://pkg.go.dev/github.com/dustin/go-humanize?tab=doc#Time
		sTime := humanize.Time(time.Unix(s.GetScheduledAt(), 0))

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(s.GetName(), s.GetEntry(), s.GetActive(), sTime, s.GetBranch())
	}

	// output the table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// wideTable is a helper function to output the
// provided schedules in a wide table format with
// a specific set of fields displayed.
func wideTable(schedules *[]library.Schedule) error {
	logrus.Debug("creating wide table for list of schedules")

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

	logrus.Trace("adding headers to wide schedule table")

	// set of schedule fields we display in a wide table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("NAME", "ENTRY", "ACTIVE", "SCHEDULED_AT", "CREATED_AT", "CREATED_BY", "UPDATED_AT", "UPDATED_BY", "BRANCH")

	// iterate through all schedules in the list
	for _, s := range *schedules {
		logrus.Tracef("adding schedule %s to schedule table", s.GetName())

		// calculate scheduled_at timestamp in human readable form
		//
		// https://pkg.go.dev/github.com/dustin/go-humanize?tab=doc#Time
		sTime := humanize.Time(time.Unix(s.GetScheduledAt(), 0))

		// calculate created_at timestamp in human readable form
		//
		// https://pkg.go.dev/github.com/dustin/go-humanize?tab=doc#Time
		cTime := humanize.Time(time.Unix(s.GetCreatedAt(), 0))

		// calculate updated_at timestamp in human readable form
		//
		// https://pkg.go.dev/github.com/dustin/go-humanize?tab=doc#Time
		uTime := humanize.Time(time.Unix(s.GetUpdatedAt(), 0))

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(s.GetName(), s.GetEntry(), s.GetActive(), sTime, cTime, s.GetCreatedBy(), uTime, s.GetUpdatedBy(), s.GetBranch())
	}

	// output the wide table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}
