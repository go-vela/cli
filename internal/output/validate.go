// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package output

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
)

// validate is a helper function to
// verify the input provided.
func validate(driver string, _input interface{}) error {
	logrus.Debugf("validating output with %s driver", driver)

	// check if the input provided is nil
	if _input == nil {
		return fmt.Errorf("empty output value provided for %s driver", driver)
	}

	// check if the value of input provided is nil
	//
	// We are using reflect here due to the nature
	// of how interfaces work in Go. It is possible
	// for _input to be a non-nil interface but the
	// underlying value to be empty or nil.
	if reflect.ValueOf(_input).IsZero() {
		return fmt.Errorf("empty output value provided for %s driver", driver)
	}

	return nil
}
