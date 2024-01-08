shell = /bin/bash

UNAME = $(shell uname)

default: help

.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

.PHONY: build
build: # Build the binary
	# Build the binary
	# ----------------
	@CGO_ENABLED=0 GOOS=linux go build -a -o build/qaparser main.go
	# Build complete +++

.PHONY: test
test: # Run a suite of tests
	# Build the binary
	# ----------------
	@go test -v ./...
	# Test complete +++