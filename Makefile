MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

export GO111MODULE = on

default: lint test

lint:
	golangci-lint run \
		--disable-all \
		--skip-files data_test.go \
		--exclude-use-default=false \
		--exclude=dot-imports \
		--exclude=package-comments \
		--exclude=unused-parameter \
		--exclude=superfluous-else \
		--enable=errcheck \
		--enable=goimports \
		--enable=ineffassign \
		--enable=revive \
		--enable=unused \
		--enable=staticcheck \
		--enable=unused \
		--enable=unconvert \
		--enable=misspell \
		--enable=prealloc \
		--enable=nakedret \
		--enable=unparam \
		./...

sec:
	gosec -quiet ./...

.PHONY: test
test:
	go test ./... -race -cover -covermode=atomic -coverprofile=unit_coverage.out
