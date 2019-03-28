#!/bin/bash

docker build -t pwrp .

docker container prune -f
docker image prune -f