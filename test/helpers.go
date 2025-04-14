// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"fmt"
	"time"

	"slices"

	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/urfave/cli/v3"
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

func TestCommand(serverURL string, action func(context.Context, *cli.Command) error, addFlags []cli.Flag) *cli.Command {
	cfgFlag := &cli.StringFlag{
		Name:  "config",
		Value: "config.yml",
	}

	for _, f := range addFlags {
		if slices.Contains(f.Names(), cfgFlag.Name) {
			cfgFlag = f.(*cli.StringFlag)
		}
	}

	cmd := &cli.Command{
		Name:   "test",
		Usage:  "Test command",
		Action: action,
		Flags: []cli.Flag{
			cfgFlag,
			&cli.StringFlag{
				Name:  "api.addr",
				Value: serverURL,
			},
			&cli.StringFlag{
				Name:  "api.token.access",
				Value: TestTokenGood,
			},
			&cli.StringFlag{
				Name:  "api.token.refresh",
				Value: "superSecretRefreshToken",
			},
		},
	}

	if len(addFlags) > 0 {
		cmd.Flags = append(cmd.Flags, addFlags...)
	}

	return cmd
}
