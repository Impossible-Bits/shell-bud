#!/usr/bin/env just --justfile

update:
    go get -u
    go mod tidy -v

build:
    go build -o ./bin/sb ./cmd/main.go

install-locally: build
    cp ./bin/sb /usr/local/bin/sb