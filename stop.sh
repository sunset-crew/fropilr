#!/bin/bash

. vars.cfg

# stop.sh
docker container start ${CONTAINER_NAME}
