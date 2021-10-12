#!/bin/bash
while :
do
	rsync -avz ./ hetzner-centos7:/root/ruchki/
	sleep 3
done