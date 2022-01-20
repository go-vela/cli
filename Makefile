
# Copyright (c) 2022 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

# capture the current date we build the application from
BUILD_DATE = $(shell date +%Y-%m-%dT%H:%M:%SZ)

# check if a git commit sha is already set
ifndef GITHUB_SHA
	# capture the current git commit sha we build the application from
	GITHUB_SHA = $(shell git rev-parse HEAD)
endif

# check if a git tag is already set
ifndef GITHUB_TAG
	# capture the current git tag we build the application from
	GITHUB_TAG = $(shell git describe --tag --abbrev=0)
endif

# create a list of linker flags for building the golang application
LD_FLAGS = -X github.com/go-vela/cli/version.Commit=${GITHUB_SHA} -X github.com/go-vela/cli/version.Date=${BUILD_DATE} -X github.com/go-vela/cli/version.Tag=${GITHUB_TAG}

# The `clean` target is intended to clean the workspace
# and prepare the local changes for submission.
#
# Usage: `make clean`
.PHONY: clean
clean: tidy vet fmt fix

# The `build` target is intended to compile
# the Go source code into a binary.
#
# Usage: `make build`
.PHONY: build
build: build-darwin build-linux build-windows

# The `build-static` target is intended to compile
# the Go source code into a statically linked binary.
#
# Usage: `make build-static`
.PHONY: build-static
build-static: build-darwin-static build-linux-static build-windows-static

# The `tidy` target is intended to clean up
# the Go module files (go.mod & go.sum).
#
# Usage: `make tidy`
.PHONY: tidy
tidy:
	@echo
	@echo "### Tidying Go module"
	@go mod tidy

# The `vet` target is intended to inspect the
# Go source code for potential issues.
#
# Usage: `make vet`
.PHONY: vet
vet:
	@echo
	@echo "### Vetting Go code"
	@go vet ./...

# The `fmt` target is intended to format the
# Go source code to meet the language standards.
#
# Usage: `make fmt`
.PHONY: fmt
fmt:
	@echo
	@echo "### Formatting Go Code"
	@go fmt ./...

# The `fix` target is intended to rewrite the
# Go source code using old APIs.
#
# Usage: `make fix`
.PHONY: fix
fix:
	@echo
	@echo "### Fixing Go Code"
	@go fix ./...

# The `test` target is intended to run
# the tests for the Go source code.
#
# Usage: `make test`
.PHONY: test
test:
	@echo
	@echo "### Testing Go Code"
	@go test -race ./...

# The `test-cover` target is intended to run
# the tests for the Go source code and then
# open the test coverage report.
#
# Usage: `make test-cover`
.PHONY: test-cover
test-cover:
	@echo
	@echo "### Creating test coverage report"
	@go test -race -covermode=atomic -coverprofile=coverage.out ./...
	@echo
	@echo "### Opening test coverage report"
	@go tool cover -html=coverage.out

# The `build-darwin` target is intended to compile the
# Go source code into a Darwin compatible (MacOS) binary.
#
# Usage: `make build-darwin`
.PHONY: build-darwin
build-darwin:
	@echo
	@echo "### Building release/darwin/amd64/vela binary"
	GOOS=darwin CGO_ENABLED=0 GOARCH=amd64 \
		go build -a \
		-ldflags '${LD_FLAGS}' \
		-o release/darwin/amd64/vela \
		github.com/go-vela/cli/cmd/vela-cli
	@echo
	@echo "### Building release/darwin/arm64/vela binary"
	GOOS=darwin CGO_ENABLED=0 GOARCH=arm64 \
		go build -a \
		-ldflags '${LD_FLAGS}' \
		-o release/darwin/arm64/vela \
		github.com/go-vela/cli/cmd/vela-cli

# The `build-linux` target is intended to compile the
# Go source code into a Linux compatible binary.
#
# Usage: `make build-linux`
.PHONY: build-linux
build-linux:
	@echo
	@echo "### Building release/linux/amd64/vela binary"
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 \
		go build -a \
		-ldflags '${LD_FLAGS}' \
		-o release/linux/amd64/vela \
		github.com/go-vela/cli/cmd/vela-cli
	@echo
	@echo "### Building release/linux/arm64/vela binary"
	GOOS=linux CGO_ENABLED=0 GOARCH=arm64 \
		go build -a \
		-ldflags '${LD_FLAGS}' \
		-o release/linux/arm64/vela \
		github.com/go-vela/cli/cmd/vela-cli
	@echo
	@echo "### Building release/linux/arm/vela binary"
	GOOS=linux CGO_ENABLED=0 GOARCH=arm \
		go build -a \
		-ldflags '${LD_FLAGS}' \
		-o release/linux/arm/vela \
		github.com/go-vela/cli/cmd/vela-cli

# The `build-windows` target is intended to compile the
# Go source code into a Windows compatible binary.
#
# Usage: `make build-windows`
.PHONY: build-windows
build-windows:
	@echo
	@echo "### Building release/windows/amd64/vela binary"
	GOOS=windows CGO_ENABLED=0 GOARCH=amd64 \
		go build -a \
		-ldflags '${LD_FLAGS}' \
		-o release/windows/amd64/vela \
		github.com/go-vela/cli/cmd/vela-cli

# The `build-darwin-static` target is intended to compile the
# Go source code into a statically linked, Darwin compatible (MacOS) binary.
#
# Usage: `make build-darwin-static`
.PHONY: build-darwin-static
build-darwin-static:
	@echo
	@echo "### Building release/darwin/amd64/vela binary"
	GOOS=darwin CGO_ENABLED=0 GOARCH=amd64 \
		go build -a \
		-ldflags '-s -w -extldflags "-static" ${LD_FLAGS}' \
		-o release/darwin/amd64/vela \
		github.com/go-vela/cli/cmd/vela-cli
	@echo
	@echo "### Building release/darwin/arm64/vela binary"
	GOOS=darwin CGO_ENABLED=0 GOARCH=amd64 \
		go build -a \
		-ldflags '-s -w -extldflags "-static" ${LD_FLAGS}' \
		-o release/darwin/arm64/vela \
		github.com/go-vela/cli/cmd/vela-cli


# The `build-linux-static` target is intended to compile the
# Go source code into a statically linked, Linux compatible binary.
#
# Usage: `make build-linux-static`
.PHONY: build-linux-static
build-linux-static:
	@echo
	@echo "### Building release/linux/amd64/vela binary"
	GOOS=linux CGO_ENABLED=0 GOARCH=amd64 \
		go build -a \
		-ldflags '-s -w -extldflags "-static" ${LD_FLAGS}' \
		-o release/linux/amd64/vela \
		github.com/go-vela/cli/cmd/vela-cli
	@echo
	@echo "### Building release/linux/arm64/vela binary"
	GOOS=linux CGO_ENABLED=0 GOARCH=arm64 \
		go build -a \
		-ldflags '-s -w -extldflags "-static" ${LD_FLAGS}' \
		-o release/linux/arm64/vela \
		github.com/go-vela/cli/cmd/vela-cli
	@echo
	@echo "### Building release/linux/arm/vela binary"
	GOOS=linux CGO_ENABLED=0 GOARCH=arm \
		go build -a \
		-ldflags '-s -w -extldflags "-static" ${LD_FLAGS}' \
		-o release/linux/arm/vela \
		github.com/go-vela/cli/cmd/vela-cli

# The `build-windows-static` target is intended to compile the
# Go source code into a statically linked, Windows compatible binary.
#
# Usage: `make build-windows-static`
.PHONY: build-windows-static
build-windows-static:
	@echo
	@echo "### Building release/windows/amd64/vela binary"
	GOOS=windows CGO_ENABLED=0 GOARCH=amd64 \
		go build -a \
		-ldflags '-s -w -extldflags "-static" ${LD_FLAGS}' \
		-o release/windows/amd64/vela \
		github.com/go-vela/cli/cmd/vela-cli

# The `check` target is intended to output all
# dependencies from the Go module that need updates.
#
# Usage: `make check`
.PHONY: check
check: check-install
	@echo
	@echo "### Checking dependencies for updates"
	@go list -u -m -json all | go-mod-outdated -update

# The `check-direct` target is intended to output direct
# dependencies from the Go module that need updates.
#
# Usage: `make check-direct`
.PHONY: check-direct
check-direct: check-install
	@echo
	@echo "### Checking direct dependencies for updates"
	@go list -u -m -json all | go-mod-outdated -direct

# The `check-full` target is intended to output
# all dependencies from the Go module.
#
# Usage: `make check-full`
.PHONY: check-full
check-full: check-install
	@echo
	@echo "### Checking all dependencies for updates"
	@go list -u -m -json all | go-mod-outdated

# The `check-install` target is intended to download
# the tool used to check dependencies from the Go module.
#
# Usage: `make check-install`
.PHONY: check-install
check-install:
	@echo
	@echo "### Installing psampaz/go-mod-outdated"
	@go get -u github.com/psampaz/go-mod-outdated

# The `bump-deps` target is intended to upgrade
# non-test dependencies for the Go module.
#
# Usage: `make bump-deps`
.PHONY: bump-deps
bump-deps: check
	@echo
	@echo "### Upgrading dependencies"
	@go get -u ./...

# The `bump-deps-full` target is intended to upgrade
# all dependencies for the Go module.
#
# Usage: `make bump-deps-full`
.PHONY: bump-deps-full
bump-deps-full: check
	@echo
	@echo "### Upgrading all dependencies"
	@go get -t -u ./...
