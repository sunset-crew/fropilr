#!/bin/bash

. vars.cfg

echo "stopping container"
docker stop ${CONTAINER_NAME}
echo "removing container"
docker rm ${CONTAINER_NAME}

if [ ! -z "$1" ] && [ "$1" = "-a" ] ; then
    docker rmi ${CONTAINER_IMAGE_NAME}:latest
else
    echo "not removing image"
    echo "use -a to delete all (container and image)"
fi
