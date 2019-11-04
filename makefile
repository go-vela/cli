# Copyright (c) 2019 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

clean:
	#################################
	######      Go clean       ######
	#################################

	# @go mod tidy
	@go vet ./...
	@go fmt ./...
	@go test ./...
	@echo "I'm kind of the only name in clean energy right now"