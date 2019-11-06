# Copyright (c) 2019 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

#!/bin/sh

set -e
set -x

# compile for all architectures
GOOS=linux   CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-X main.version=${VELA_TAG}" -o release/linux/amd64/vela   github.com/go-vela/cli
GOOS=linux   CGO_ENABLED=0 GOARCH=arm64 go build -ldflags "-X main.version=${VELA_TAG}" -o release/linux/arm64/vela   github.com/go-vela/cli
GOOS=linux   CGO_ENABLED=0 GOARCH=arm   go build -ldflags "-X main.version=${VELA_TAG}" -o release/linux/arm/vela     github.com/go-vela/cli
GOOS=windows CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-X main.version=${VELA_TAG}" -o release/windows/amd64/vela github.com/go-vela/cli
GOOS=darwin  CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-X main.version=${VELA_TAG}" -o release/darwin/amd64/vela  github.com/go-vela/cli
