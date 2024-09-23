# SPDX-License-Identifier: Apache-2.0

#################################################
#  docker build -t target/vela-cli:latest .     #
#################################################

FROM alpine:3.20.3@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d

RUN apk add --update --no-cache ca-certificates

COPY release/linux/amd64/vela /bin/vela

CMD ["/bin/vela"]
