// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package secret

import (
	"fmt"
	"strings"

	"github.com/go-vela/cli/internal/output"

	"github.com/go-vela/types/constants"
	"github.com/go-vela/types/library"

	"github.com/gosuri/uitable"

	"github.com/sirupsen/logrus"
)

// table is a helper function to output the
// provided secrets in a table format with
// a specific set of fields displayed.
func table(secrets *[]library.Secret) error {
	logrus.Debug("creating table for list of secrets")

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

	logrus.Trace("adding headers to secret table")

	// set of secret fields we display in a table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("NAME", "ORG", "TYPE", "KEY")

	// iterate through all secrets in the list
	for _, s := range *secrets {
		logrus.Tracef("adding secret %s to secret table", s.GetName())

		// calculate the key for the secret
		//
		// nolint: gosec // ignore memory aliasing
		k := key(&s)

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(s.GetName(), s.GetOrg(), s.GetType(), k)
	}

	// output the table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// wideTable is a helper function to output the
// provided secrets in a wide table format with
// a specific set of fields displayed.
func wideTable(secrets *[]library.Secret) error {
	logrus.Debug("creating wide table for list of secrets")

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

	logrus.Trace("adding headers to wide secret table")

	// set of secret fields we display in a wide table
	//
	// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
	table.AddRow("NAME", "ORG", "TYPE", "KEY", "EVENTS", "IMAGES")

	// iterate through all secrets in the list
	for _, s := range *secrets {
		logrus.Tracef("adding secret %s to wide secret table", s.GetName())

		// capture list of events for secret
		e := strings.Join(s.GetEvents(), ",")

		// capture list of images for secret
		i := strings.Join(s.GetImages(), ",")

		// calculate the key for the secret
		//
		// nolint: gosec // ignore memory aliasing
		k := key(&s)

		// add a row to the table with the specified values
		//
		// https://pkg.go.dev/github.com/gosuri/uitable?tab=doc#Table.AddRow
		table.AddRow(s.GetName(), s.GetOrg(), s.GetType(), k, e, i)
	}

	// output the wide table in stdout format
	//
	// https://pkg.go.dev/github.com/go-vela/cli/internal/output?tab=doc#Stdout
	return output.Stdout(table)
}

// key is a helper function to calculate the full
// path to a secret in the storage backend.
func key(s *library.Secret) string {
	switch s.GetType() {
	case constants.SecretShared:
		return fmt.Sprintf("%s/%s/%s", s.GetOrg(), s.GetTeam(), s.GetName())
	case constants.SecretOrg:
		return fmt.Sprintf("%s/%s", s.GetOrg(), s.GetName())
	case constants.SecretRepo:
		fallthrough
	default:
		return fmt.Sprintf("%s/%s/%s", s.GetOrg(), s.GetRepo(), s.GetName())
	}
}
