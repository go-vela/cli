// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package deployment

import (
	"sort"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/types/library"

	"github.com/gosuri/uitable"
)

// table is a helper function to output the
// provided deployments in a table format with
// a specific set of fields displayed.
func table(deployments *[]library.Deployment) error {
	// create new table
	table := uitable.New()

	// set column width for table to 50
	table.MaxColWidth = 50

	// ensure the table is always wrapped
	table.Wrap = true

	// set of deployment fields we display in a table
	table.AddRow("ID", "TASK", "USER", "REF", "TARGET")

	// iterate through all deployments in the list
	for _, d := range reverse(*deployments) {
		// add a row to the table with the specified values
		table.AddRow(d.GetID(), d.GetTask(), d.GetUser(), d.GetRef(), d.GetTarget())
	}

	// output the table in stdout format
	err := output.Stdout(table)
	if err != nil {
		return err
	}

	return nil
}

// wideTable is a helper function to output the
// provided deployments in a wide table format with
// a specific set of fields displayed.
func wideTable(deployments *[]library.Deployment) error {
	// create new wide table
	table := uitable.New()

	// set column width for wide table to 200
	table.MaxColWidth = 200

	// ensure the wide table is always wrapped
	table.Wrap = true

	// set of deployment fields we display in a wide table
	table.AddRow("ID", "TASK", "USER", "REF", "TARGET", "COMMIT", "DESCRIPTION")

	// iterate through all deployments in the list
	for _, d := range reverse(*deployments) {
		// add a row to the table with the specified values
		table.AddRow(d.GetID(), d.GetTask(), d.GetUser(), d.GetRef(), d.GetTarget(), d.GetCommit(), d.GetDescription())
	}

	// output the wide table in stdout format
	err := output.Stdout(table)
	if err != nil {
		return err
	}

	return nil
}

// reverse is a helper function to sort the deployments
// based off the deployment number and then flip the
// order they get displayed in.
func reverse(d []library.Deployment) []library.Deployment {
	sort.SliceStable(d, func(i, j int) bool {
		return d[i].GetID() < d[j].GetID()
	})

	return d
}
