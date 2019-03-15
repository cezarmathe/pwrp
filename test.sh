#!/bin/bash

export PWRP_DEBUG=true

go run main.go --config .local/pwrp.toml $@

