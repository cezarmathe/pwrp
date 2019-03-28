#!/bin/bash

docker build -t cezarmathe:pwrp .

docker run cezarmathe:pwrp --env PWRP_DEBUG=$1


