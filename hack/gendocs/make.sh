#!/usr/bin/env bash

pushd $GOPATH/src/voyagermesh.dev/hello-grpc/hack/gendocs
go run main.go
popd
