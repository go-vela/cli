// SPDX-License-Identifier: Apache-2.0

package worker

import (
	"time"

	"github.com/go-vela/cli/internal/output"
	api "github.com/go-vela/server/api/types"

	"github.com/dustin/go-humanize"
	"github.com/gosuri/uitable"
	"github.com/sirupsen/logrus"
)

// table is a helper function to output the
// provided workers in a table format with
// a specific set of fields displayed.
func table(workers *[]api.Worker) error {
	logrus.Debug("creating table for list of workers")

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

	logrus.Trace("adding headers to worker table")

	// set of worker fields we display in a table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("HOSTNAME", "ACTIVE", "ROUTES", "LAST_CHECKED_IN")

	// iterate through all workers in the list
	for _, w := range *workers {
		logrus.Tracef("adding worker %s to worker table", w.GetHostname())

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(w.GetHostname(), w.GetActive(), w.GetRoutes(), w.GetLastCheckedIn())
	}

	// output the table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// wideTable is a helper function to output the
// provided workers in a wide table format with
// a specific set of fields displayed.
func wideTable(workers *[]api.Worker) error {
	logrus.Debug("creating wide table for list of workers")

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

	logrus.Trace("adding headers to wide worker table")

	// set of worker fields we display in a wide table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("HOSTNAME", "ADDRESS", "ACTIVE", "ROUTES", "LAST_CHECKED_IN", "BUILD_LIMIT")

	// iterate through all workers in the list
	for _, w := range *workers {
		logrus.Tracef("adding worker %s to wide worker table", w.GetHostname())

		// calculate last checked in timestamp in human readable form
		//
		// https://pkg.go.dev/github.com/dustin/go-humanize?tab=doc#Time
		c := humanize.Time(time.Unix(w.GetLastCheckedIn(), 0))

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(w.GetHostname(), w.GetAddress(), w.GetActive(), w.GetRoutes(), c, w.GetBuildLimit())
	}

	// output the wide table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}
