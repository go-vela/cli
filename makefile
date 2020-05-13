# Copyright (c) 2020 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

clean:
	#################################
	######      Go clean       ######
	#################################

	@go mod tidy
	@go vet ./...
	@go fmt ./...
	@go test ./...
	@echo "I'm kind of the only name in clean energy right now"

build:
	#################################
	###### Build Golang Binary ######
	#################################

	GOOS=linux   CGO_ENABLED=0 GOARCH=amd64 go build -o release/linux/amd64/vela   github.com/go-vela/cli/cmd/vela-cli
	GOOS=linux   CGO_ENABLED=0 GOARCH=arm64 go build -o release/linux/arm64/vela   github.com/go-vela/cli/cmd/vela-cli
	GOOS=linux   CGO_ENABLED=0 GOARCH=arm   go build -o release/linux/arm/vela     github.com/go-vela/cli/cmd/vela-cli
	GOOS=windows CGO_ENABLED=0 GOARCH=amd64 go build -o release/windows/amd64/vela github.com/go-vela/cli/cmd/vela-cli
	GOOS=darwin  CGO_ENABLED=0 GOARCH=amd64 go build -o release/darwin/amd64/vela  github.com/go-vela/cli/cmd/vela-cli
