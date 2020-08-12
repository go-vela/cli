# Copyright (c) 2020 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

############################################################
#  docker build --target certs -t target/vela-cli:certs .  #
############################################################

FROM alpine as certs

RUN apk add --update --no-cache ca-certificates

#################################################
#  docker build -t target/vela-cli:latest .     #
#################################################

FROM alpine

COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY release/linux/amd64/vela /bin/vela

CMD ["/bin/vela"]
