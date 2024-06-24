# SPDX-License-Identifier: Apache-2.0

#################################################
#  docker build -t target/vela-cli:latest .     #
#################################################

FROM alpine:3.20.1@sha256:b89d9c93e9ed3597455c90a0b88a8bbb5cb7188438f70953fede212a0c4394e0

RUN apk add --update --no-cache ca-certificates

COPY release/linux/amd64/vela /bin/vela

CMD ["/bin/vela"]
