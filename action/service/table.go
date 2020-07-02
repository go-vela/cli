// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package service

import (
	"sort"
	"strings"
	"time"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"

	"github.com/dustin/go-humanize"
	"github.com/gosuri/uitable"
)

// table is a helper function to output the
// provided services in a table format with
// a specific set of fields displayed.
func table(services *[]library.Service) error {
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

	// set of service fields we display in a table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("NUMBER", "NAME", "STATUS", "DURATION")

	// iterate through all services in the list
	for _, s := range reverse(*services) {
		// calculate duration based off the service timestamps
		d := duration(&s)

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(s.GetNumber(), s.GetName(), s.GetStatus(), d)
	}

	// output the table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// wideTable is a helper function to output the
// provided services in a wide table format with
// a specific set of fields displayed.
func wideTable(services *[]library.Service) error {
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

	// set of service fields we display in a wide table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("NUMBER", "NAME", "STATUS", "DURATION", "CREATED", "FINISHED")

	// iterate through all services in the list
	for _, s := range reverse(*services) {
		// calculate duration based off the service timestamps
		d := duration(&s)

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
		table.AddRow(s.GetNumber(), s.GetName(), s.GetStatus(), d, c, f)
	}

	// output the wide table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// duration is a helper function to calculate
// the total duration a service ran for in a
// more consumable, human readable format.
func duration(s *library.Service) string {
	// check if service is in a pending or running state
	if strings.EqualFold(s.GetStatus(), constants.StatusPending) ||
		strings.EqualFold(s.GetStatus(), constants.StatusRunning) {
		// return a default value to display the service is not complete
		return "..."
	}

	// capture finished unix timestamp from service
	f := time.Unix(s.GetFinished(), 0)
	// capture started unix timestamp from service
	st := time.Unix(s.GetStarted(), 0)

	// get the duration by subtracting the service
	// started unix timestamp from the service finished
	// unix timestamp.
	d := f.Sub(st)

	// return duration in a human readable form
	return d.String()
}

// reverse is a helper function to sort the services
// based off the service number and then flip the
// order they get displayed in.
func reverse(s []library.Service) []library.Service {
	// sort the list of services based off the service number
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].GetNumber() < s[j].GetNumber()
	})

	return s
}
