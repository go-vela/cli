# SPDX-License-Identifier: Apache-2.0

#################################################
#  docker build -t target/vela-cli:latest .     #
#################################################

FROM alpine:3.24.3@sha256:be2656061b536c1883c4034790e21d0a09c5aaaf51968b89368a3c0371ea2c9a

RUN apk add --update --no-cache ca-certificates

COPY release/linux/amd64/vela /bin/vela

CMD ["/bin/vela"]
