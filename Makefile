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

update-deps:
	go get -u go.aporeto.io/regolithe@master
	go get -u github.com/smartystreets/goconvey@latest
	go get -u go.uber.org/zap@latest
	go get -u golang.org/x/sync@latest
	go get -u golang.org/x/text@latest
	go get -u golang.org/x/tools@latest

	go mod tidy
