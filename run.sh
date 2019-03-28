#!/bin/bash

docker build -t pwrp .

docker run pwrp -e PWRP_DEBUG=$1


