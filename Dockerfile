# SPDX-License-Identifier: Apache-2.0

#################################################
#  docker build -t target/vela-cli:latest .     #
#################################################

FROM alpine:3.19.0@sha256:51b67269f354137895d43f3b3d810bfacd3945438e94dc5ac55fdac340352f48

RUN apk add --update --no-cache ca-certificates

COPY release/linux/amd64/vela /bin/vela

CMD ["/bin/vela"]
