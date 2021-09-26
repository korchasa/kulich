#!/usr/bin/env bash

docker build --tag centos7 . -f ./centos7.Dockerfile
docker run -it -v "$(pwd)":/work centos7 bash
