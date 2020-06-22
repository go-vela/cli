// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
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
)

// table is a helper function to output the
// provided secrets in a table format with
// a specific set of fields displayed.
func table(secrets *[]library.Secret) error {
	// create new table
	table := uitable.New()

	// set column width for table to 50
	table.MaxColWidth = 50

	// ensure the table is always wrapped
	table.Wrap = true

	// set of secret fields we display in a table
	table.AddRow("NAME", "ORG", "TYPE", "KEY")

	// iterate through all secrets in the list
	for _, s := range *secrets {
		// calculate the key for the secret
		k := key(&s)

		// add a row to the table with the specified values
		table.AddRow(s.GetName(), s.GetOrg(), s.GetType(), k)
	}

	// output the table in default format
	err := output.Default(table)
	if err != nil {
		return err
	}

	return nil
}

// wideTable is a helper function to output the
// provided secrets in a wide table format with
// a specific set of fields displayed.
func wideTable(secrets *[]library.Secret) error {
	// create new wide table
	table := uitable.New()

	// set column width for wide table to 200
	table.MaxColWidth = 200

	// ensure the wide table is always wrapped
	table.Wrap = true

	// set of secret fields we display in a wide table
	table.AddRow("NAME", "ORG", "TYPE", "KEY", "EVENTS", "IMAGES")

	// iterate through all secrets in the list
	for _, s := range *secrets {
		// capture list of events for secret
		e := strings.Join(s.GetEvents(), ",")

		// capture list of images for secret
		i := strings.Join(s.GetImages(), ",")

		// calculate the key for the secret
		k := key(&s)

		// add a row to the table with the specified values
		table.AddRow(s.GetName(), s.GetOrg(), s.GetType(), k, e, i)
	}

	// output the wide table in default format
	err := output.Default(table)
	if err != nil {
		return err
	}

	return nil
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
