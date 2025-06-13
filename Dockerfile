# SPDX-License-Identifier: Apache-2.0

#################################################
#  docker build -t target/vela-cli:latest .     #
#################################################

FROM alpine:3.22.0@sha256:8a1f59ffb675680d47db6337b49d22281a139e9d709335b492be023728e11715

RUN apk add --update --no-cache ca-certificates

COPY release/linux/amd64/vela /bin/vela

CMD ["/bin/vela"]
