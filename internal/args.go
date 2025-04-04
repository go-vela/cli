// SPDX-License-Identifier: Apache-2.0

package internal

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v3"
)

// ProcessArgs attempts to get the command line
// arguments, grab the first value, and set the
// resource to that value in the context.
func ProcessArgs(c *cli.Context, resource string, expect string) error {
	args := c.Args()

	val := args.First()
	if val == "" {
		return nil
	}

	if expect == "int" {
		_, err := strconv.Atoi(val)
		if err != nil {
			retErr := fmt.Errorf("invalid type for %s: expect %s", resource, expect)
			return retErr
		}
	}

	err := c.Set(resource, val)
	if err != nil {
		return err
	}

	return nil
}
