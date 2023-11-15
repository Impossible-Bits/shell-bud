#!/usr/bin/env just --justfile

update:
    go get -u
    go mod tidy -v

build:
    go build -o ./bin/shm ./cmd/main.go

install-locally: build
    cp ./bin/shm /usr/local/bin/shm