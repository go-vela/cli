// SPDX-License-Identifier: Apache-2.0

package test

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// expose some pre-computed test tokens.
var (
	TestTokenGood    = makeSampleToken(jwt.MapClaims{"exp": float64(time.Now().Unix() + 100)})
	TestTokenExpired = makeSampleToken(jwt.MapClaims{"exp": float64(time.Now().Unix() - 100)})
)

// makeSampleToken is a helper to create test tokens
// with the given claims.
func makeSampleToken(c jwt.Claims) string {
	// create a new token
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, c)

	// get the signing string (header + claims)
	s, e := t.SigningString()

	if e != nil {
		return ""
	}

	// add bogus signature
	s = fmt.Sprintf("%s.abcdef", s)

	return s
}
