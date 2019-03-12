#!/bin/bash

if [[ -z "$1" ]]; then
    export DEBUG_LOG_LEVEL=true
fi

go build

cp pppi .local/test/pppi

rm pppi

cd .local/test

./pppi --cfg "$(pwd)/index_config.toml"

