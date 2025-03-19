#!/bin/sh

source ./.env

pushd emulator
if [[ "$1" == "down" ]]; then
    docker compose down
else
    docker compose up --build --wait -t 90
fi
popd