# Copyright (c) 2020 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

#!/bin/sh

set -e
set -x

# capture tag version from reference
tag=$(echo ${GITHUB_REF} | cut -d / -f 3)

# compile for all architectures
GOOS=linux   CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-X main.version=${tag}" -o release/linux/amd64/vela   github.com/go-vela/cli/cmd/vela-cli
GOOS=linux   CGO_ENABLED=0 GOARCH=arm64 go build -ldflags "-X main.version=${tag}" -o release/linux/arm64/vela   github.com/go-vela/cli/cmd/vela-cli
GOOS=linux   CGO_ENABLED=0 GOARCH=arm   go build -ldflags "-X main.version=${tag}" -o release/linux/arm/vela     github.com/go-vela/cli/cmd/vela-cli
GOOS=windows CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-X main.version=${tag}" -o release/windows/amd64/vela github.com/go-vela/cli/cmd/vela-cli
GOOS=darwin  CGO_ENABLED=0 GOARCH=amd64 go build -ldflags "-X main.version=${tag}" -o release/darwin/amd64/vela  github.com/go-vela/cli/cmd/vela-cli

# tar binary files prior to upload
tar -cvzf release/vela_linux_amd64.tar.gz   -C release/linux/amd64   vela
tar -cvzf release/vela_linux_arm64.tar.gz   -C release/linux/arm64   vela
tar -cvzf release/vela_linux_arm.tar.gz     -C release/linux/arm     vela
tar -cvzf release/vela_windows_amd64.tar.gz -C release/windows/amd64 vela
tar -cvzf release/vela_darwin_amd64.tar.gz  -C release/darwin/amd64  vela

# generate shas for tar files
sha256sum release/*.tar.gz > release/vela_checksums.txt
