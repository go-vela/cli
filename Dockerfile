# Copyright (c) 2021 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

#################################################
#  docker build -t target/vela-cli:latest .     #
#################################################

FROM alpine:20210730

RUN apk add --update --no-cache ca-certificates

COPY release/linux/amd64/vela /bin/vela

CMD ["/bin/vela"]
