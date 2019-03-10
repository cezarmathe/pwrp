#!/bin/bash

go build

cp pppi .local/test/pppi

rm pppi

cd .local/test

./pppi -i "$(pwd)/index_config.toml"

