#!/bin/sh

. ./emulate.sh

pushd services/crowdify/integration
go test -count=1 ./...
popd