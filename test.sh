#!/usr/bin/env bash
set -ex

docker build --tag ruchki-lint -f ./lint.Dockerfile .
docker run ruchki-lint

docker build --tag ruchki-centos7 -f ./centos7.Dockerfile .
docker run ruchki-centos7
