// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package step

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
// provided steps in a table format with
// a specific set of fields displayed.
func table(steps *[]library.Step) error {
	// create new table
	table := uitable.New()

	// set column width for table to 50
	table.MaxColWidth = 50

	// ensure the table is always wrapped
	table.Wrap = true

	// set of step fields we display in a table
	table.AddRow("NUMBER", "NAME", "STATUS", "DURATION")

	// iterate through all steps in the list
	for _, s := range reverse(*steps) {
		// calculate duration based off the step timestamps
		d := duration(&s)

		// add a row to the table with the specified values
		table.AddRow(s.GetNumber(), s.GetName(), s.GetStatus(), d)
	}

	// output the table in default format
	err := output.Default(table)
	if err != nil {
		return err
	}

	return nil
}

// wideTable is a helper function to output the
// provided steps in a wide table format with
// a specific set of fields displayed.
func wideTable(steps *[]library.Step) error {
	// create new wide table
	table := uitable.New()

	// set column width for wide table to 200
	table.MaxColWidth = 200

	// ensure the wide table is always wrapped
	table.Wrap = true

	// set of step fields we display in a wide table
	table.AddRow("NUMBER", "NAME", "STATUS", "DURATION", "CREATED", "FINISHED")

	// iterate through all steps in the list
	for _, s := range reverse(*steps) {
		// calculate duration based off the step timestamps
		d := duration(&s)

		// calculate created timestamp in human readable form
		c := humanize.Time(time.Unix(s.GetCreated(), 0))

		// calculate finished timestamp in human readable form
		f := humanize.Time(time.Unix(s.GetFinished(), 0))

		// add a row to the table with the specified values
		table.AddRow(s.GetNumber(), s.GetName(), s.GetStatus(), d, c, f)
	}

	// output the wide table in default format
	err := output.Default(table)
	if err != nil {
		return err
	}

	return nil
}

// duration is a helper function to calculate
// the total duration a step ran for in a
// more consumable, human readable format.
func duration(s *library.Step) string {
	// check if step is in a pending or running state
	if strings.EqualFold(s.GetStatus(), constants.StatusPending) ||
		strings.EqualFold(s.GetStatus(), constants.StatusRunning) {
		// return a default value to display the step is not complete
		return "..."
	}

	// capture finished unix timestamp from step
	f := time.Unix(s.GetFinished(), 0)
	// capture started unix timestamp from step
	st := time.Unix(s.GetStarted(), 0)

	// get the duration by subtracting the step
	// started unix timestamp from the step finished
	// unix timestamp.
	d := f.Sub(st)

	// return duration in a human readable form
	return d.String()
}

// reverse is a helper function to sort the steps
// based off the step number and then flip the
// order they get displayed in.
func reverse(s []library.Step) []library.Step {
	sort.SliceStable(s, func(i, j int) bool {
		return s[i].GetNumber() < s[j].GetNumber()
	})

	return s
}
