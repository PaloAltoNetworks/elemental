MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

export GO111MODULE = on

default: lint test

lint:
	golangci-lint run ./...

sec:
	gosec -quiet ./...

.PHONY: test
test:
	go test ./... -race -cover -covermode=atomic -coverprofile=unit_coverage.out

update-deps:
	go get -u go.aporeto.io/regolithe@master
	go get -u github.com/smartystreets/goconvey@latest
	go get -u go.uber.org/zap@latest
	go get -u golang.org/x/sync@latest
	go get -u golang.org/x/text@latest
	go get -u golang.org/x/tools@latest

	go mod tidy
