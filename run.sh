#!/bin/bash

. vars.cfg

docker run -d -t --hostname ${DHOSTNAME} --name ${CONTAINER_NAME} \
  -p 9999:9999 \
  ${CONTAINER_IMAGE_NAME}:latest

#~ docker run -d -t --hostname ${DHOSTNAME} --name ${CONTAINER_NAME} \
  #~ --security-opt seccomp=unconfined \
  #~ --tmpfs /run --tmpfs /run/lock --v /sys/fs/cgroup:/sys/fs/cgroup:ro \
  #~ -p 587:587 -p 993:993 -p 8081:8081 \
  #~ -p 8081:8081 \
  #~ ${CONTAINER_IMAGE_NAME}:latest /lib/systemd/systemd
