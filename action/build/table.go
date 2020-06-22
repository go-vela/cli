// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package build

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
// provided builds in a table format with
// a specific set of fields displayed.
func table(builds *[]library.Build) error {
	// create new table
	table := uitable.New()

	// set column width for table to 50
	table.MaxColWidth = 50

	// ensure the table is always wrapped
	table.Wrap = true

	// set of build fields we display in a table
	table.AddRow("NUMBER", "STATUS", "EVENT", "BRANCH", "DURATION")

	// iterate through all builds in the list
	for _, b := range reverse(*builds) {
		// calculate duration based off the build timestamps
		d := duration(&b)

		// add a row to the table with the specified values
		table.AddRow(b.GetNumber(), b.GetStatus(), b.GetEvent(), b.GetBranch(), d)
	}

	// output the table in default format
	err := output.Default(table)
	if err != nil {
		return err
	}

	return nil
}

// wideTable is a helper function to output the
// provided builds in a wide table format with
// a specific set of fields displayed.
func wideTable(builds *[]library.Build) error {
	// create new wide table
	table := uitable.New()

	// set column width for wide table to 200
	table.MaxColWidth = 200

	// ensure the wide table is always wrapped
	table.Wrap = true

	// set of build fields we display in a wide table
	table.AddRow("NUMBER", "STATUS", "EVENT", "BRANCH", "COMMIT", "DURATION", "CREATED", "FINISHED", "AUTHOR")

	// iterate through all builds in the list
	for _, b := range reverse(*builds) {
		// calculate duration based off the build timestamps
		d := duration(&b)

		// calculate created timestamp in human readable form
		c := humanize.Time(time.Unix(b.GetCreated(), 0))

		// calculate finished timestamp in human readable form
		f := humanize.Time(time.Unix(b.GetFinished(), 0))

		// add a row to the table with the specified values
		table.AddRow(b.GetNumber(), b.GetStatus(), b.GetEvent(), b.GetBranch(), b.GetCommit(), d, c, f, b.GetAuthor())
	}

	// output the wide table in default format
	err := output.Default(table)
	if err != nil {
		return err
	}

	return nil
}

// duration is a helper function to calculate
// the total duration a build ran for in a
// more consumable, human readable format.
func duration(b *library.Build) string {
	// check if build is in a pending or running state
	if strings.EqualFold(b.GetStatus(), constants.StatusPending) ||
		strings.EqualFold(b.GetStatus(), constants.StatusRunning) {
		// return a default value to display the build is not complete
		return "..."
	}

	// capture finished unix timestamp from build
	f := time.Unix(b.GetFinished(), 0)
	// capture started unix timestamp from build
	s := time.Unix(b.GetStarted(), 0)

	// get the duration by subtracting the build
	// started unix timestamp from the build finished
	// unix timestamp.
	d := f.Sub(s)

	// return duration in a human readable form
	return d.String()
}

// reverse is a helper function to sort the builds
// based off the build number and then flip the
// order they get displayed in.
func reverse(b []library.Build) []library.Build {
	sort.SliceStable(b, func(i, j int) bool {
		return b[i].GetNumber() < b[j].GetNumber()
	})

	return b
}
