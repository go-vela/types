# Copyright (c) 2020 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

.PHONY: restart
restart: down up

.PHONY: up
up: build compose-up

.PHONY: down
down: compose-down

.PHONY: rebuild
rebuild: build compose-up

.PHONY: clean
clean:
	#################################
	######      Go clean       ######
	#################################

	@go mod tidy
	@go vet ./...
	@go fmt ./...
	@echo "I'm kind of the only name in clean energy right now"
