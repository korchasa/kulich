#!/usr/bin/env bash

docker build --progress plain --no-cache --tag ruchki-centos7-dev . -f ./centos7.Dockerfile
docker run --detach --rm -v "$(pwd)":/work --name ruchki-centos7-dev ruchki-centos7-dev
docker exec -it ruchki-centos7-dev bash

