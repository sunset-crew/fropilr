#!/bin/bash

. vars.cfg

make build
# prebuild
docker build -t ${CONTAINER_IMAGE_NAME}:latest .
# postbuild
