# ------------------------------------------------
# Copyright (C) 2016 Aporeto Inc.
#
# File  : Makefile
#
# Author: alex@aporeto.com, antoine@aporeto.com
# Date  : 2016-03-8
#
# ------------------------------------------------

## configure this throught environment variables
PROJECT_OWNER?=github.com/aporeto-inc
PROJECT_NAME?=my-super-project
DOMINGO_DOCKER_TAG?=latest
DOMINGO_DOCKER_REPO=gcr.io/aporetodev
GITHUB_TOKEN?=

######################################################################
######################################################################

export ROOT_DIR?=$(PWD)

MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

APOMOCK_FILE            := .apomock
APOMOCK_PACKAGES        := $(shell if [ -f $(APOMOCK_FILE) ]; then cat $(APOMOCK_FILE); fi)
NOVENDOR                := $(shell glide novendor)
MANAGED_DIRS            := $(sort $(dir $(wildcard */Makefile)))
MOCK_DIRS               := $(sort $(dir $(wildcard */.apomock)))
NOTEST_DIRS             := $(MANAGED_DIRS)
NOTEST_DIRS             := $(addsuffix ...,$(NOTEST_DIRS))
NOTEST_DIRS             := $(addprefix ./,$(NOTEST_DIRS))
TEST_DIRS               := $(filter-out $(NOTEST_DIRS),$(NOVENDOR))
GO_SRCS                 := $(wildcard *.go)

## Update

domingo_update:
	@echo "# Updating Domingo..."
	@echo "REMINDER: you need to export GITHUB_TOKEN for this to work"
	curl --fail -o domingo.mk -H "Cache-Control: no-cache" -H "Authorization: token $(GITHUB_TOKEN)" https://raw.githubusercontent.com/aporeto-inc/domingo/master/domingo.mk
	@echo "domingo.mk updated!"


## initialization

domingo_init:
	@if [ -f glide.yaml ]; then glide install; else go get ./...; fi

## Testing

domingo_goconvey:
	make domingo_lint domingo_init_apomock
	goconvey .
	make domingo_deinit_apomock

domingo_test:
	@$(foreach dir,$(MANAGED_DIRS),pushd ${dir} > /dev/null && make domingo_test && popd > /dev/null;)
	@if [ -f $(APOMOCK_FILE) ]; then make domingo_init_apomock; fi
	@if [ "$(GO_SRCS)" != "" ]; then go test -race -cover $(TEST_DIRS) || exit 1; else echo "# Skipped as no go sources found"; fi
	@if [ -f $(APOMOCK_FILE) ]; then make domingo_deinit_apomock; fi


domingo_init_apomock:
	@make domingo_save_vendor
	@kennebec --package="$(APOMOCK_PACKAGES)" --output-dir=vendor -v=4 -logtostderr=true >> /dev/null 2>&1

domingo_deinit_apomock:
	@make domingo_restore_vendor

domingo_save_vendor:
	@if [ -d vendor ]; then cp -a vendor vendor.lock; fi

domingo_restore_vendor:
	@if [ -d vendor.lock ]; then rm -rf vendor && mv vendor.lock vendor; else rm -rf vendor; fi

container:
	make build_linux
	cd docker && docker build -t $(DOMINGO_DOCKER_REPO)/$(PROJECT_NAME):$(DOMINGO_DOCKER_TAG) .
