// SPDX-License-Identifier: Apache-2.0

package service

import (
	"sort"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gosuri/uitable"
	"github.com/sirupsen/logrus"

	"github.com/go-vela/cli/internal/output"
	api "github.com/go-vela/server/api/types"
)

// table is a helper function to output the
// provided services in a table format with
// a specific set of fields displayed.
func table(services *[]api.Service) error {
	logrus.Debug("creating table for list of services")

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

	logrus.Trace("adding headers to service table")

	// set of service fields we display in a table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("NUMBER", "NAME", "STATUS", "DURATION")

	// iterate through all services in the list
	for _, s := range reverse(*services) {
		logrus.Tracef("adding service %d to service table", s.GetNumber())

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(s.GetNumber(), s.GetName(), s.GetStatus(), s.Duration())
	}

	// output the table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// wideTable is a helper function to output the
// provided services in a wide table format with
// a specific set of fields displayed.
func wideTable(services *[]api.Service) error {
	logrus.Debug("creating wide table for list of services")

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

	logrus.Trace("adding headers to wide service table")

	// set of service fields we display in a wide table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("NUMBER", "NAME", "STATUS", "DURATION", "CREATED", "FINISHED")

	// iterate through all services in the list
	for _, s := range reverse(*services) {
		logrus.Tracef("adding service %d to wide service table", s.GetNumber())

		// calculate created timestamp in human readable form
		//
		// https://pkg.go.dev/github.com/dustin/go-humanize?tab=doc#Time
		c := humanize.Time(time.Unix(s.GetCreated(), 0))

		// calculate finished timestamp in human readable form
		//
		// https://pkg.go.dev/github.com/dustin/go-humanize?tab=doc#Time
		f := humanize.Time(time.Unix(s.GetFinished(), 0))

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(s.GetNumber(), s.GetName(), s.GetStatus(), s.Duration(), c, f)
	}

	// output the wide table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// reverse is a helper function to sort the services
// based off the service number and then flip the
// order they get displayed in.
func reverse(s []api.Service) []api.Service {
	// sort the list of services based off the service number
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].GetNumber() < s[j].GetNumber()
	})

	return s
}
