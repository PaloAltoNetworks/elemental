MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

export GO111MODULE = on

default: lint test sec

lint:
	golangci-lint run \
		--disable-all \
		--skip-files data_test.go \
		--exclude-use-default=false \
		--enable=errcheck \
		--enable=goimports \
		--enable=ineffassign \
		--enable=golint \
		--enable=unused \
		--enable=structcheck \
		--enable=staticcheck \
		--enable=varcheck \
		--enable=deadcode \
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
	go test ./... -race -cover -covermode=atomic -coverprofile=unit_coverage.cov
