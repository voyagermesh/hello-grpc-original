#!/usr/bin/env bash

pushd $GOPATH/src/github.com/appscode/hello-grpc/hack/gendocs
go run main.go
popd
