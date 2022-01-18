# Copyright (c) 2022 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

#!/bin/sh

set -e
set -x

# compile for all architectures
make build

# tar binary files prior to upload
tar -cvzf release/vela_linux_amd64.tar.gz   -C release/linux/amd64   vela
tar -cvzf release/vela_linux_arm64.tar.gz   -C release/linux/arm64   vela
tar -cvzf release/vela_linux_arm.tar.gz     -C release/linux/arm     vela
tar -cvzf release/vela_windows_amd64.tar.gz -C release/windows/amd64 vela
tar -cvzf release/vela_darwin_amd64.tar.gz  -C release/darwin/amd64  vela
tar -cvzf release/vela_darwin_arm64.tar.gz  -C release/darwin/arm64  vela

# generate shas for tar files
sha256sum release/*.tar.gz > release/vela_checksums.txt
