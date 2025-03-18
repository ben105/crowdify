#!/bin/sh

pushd emulator
if [[ "$1" -eq "down" ]]; then
    docker compose down
else
    docker compose up --wait -t 90
fi
popd