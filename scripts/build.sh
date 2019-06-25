#!/bin/bash

export GO111MODULE=on

go get golang.org/x/tools/cmd/cover
go get github.com/mattn/goveralls


go build -v ./...
go test -v -covermode=count -coverprofile=coverage.out
$HOME/gopath/bin/goveralls -coverprofile=coverage.out -service=travis-ci -repotoken $COVERALLS_TOKEN