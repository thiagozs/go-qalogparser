shell = /bin/bash

UNAME = $(shell uname)

default: help

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: mocks-interfaces
mocks-interfaces: # Generate mocks and interfaces
	# Generate mocks
	# --------------
	@go generate ./...
	# Mocks complete +++

.PHONY: install-dep
install-dep: # Install dependencies
	# Install dependencies
	# --------------------
	# install ifacemaker...
	@go install github.com/vburenin/ifacemaker@latest
	# install mockgen...
	@go install go.uber.org/mock/mockgen@latest
	# install cobra-cli...
	@go install github.com/spf13/cobra-cli@latest
	# Install deps complete +++