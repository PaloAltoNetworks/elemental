#!/bin/bash

gometalinter \
    --exclude bindata.go \
    --exclude vendor \
    --vendor \
    --disable-all \
    --enable vet \
    --enable vetshadow \
    --enable golint \
    --enable ineffassign \
    --enable goconst \
    --enable errcheck \
    --enable varcheck \
    --enable structcheck \
    --enable gosimple \
    --enable misspell \
    --enable deadcode \
    --enable staticcheck \
    --deadline 5m \
    --tests ./...

go test -race -cover ./...
