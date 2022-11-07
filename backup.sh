#!/bin/bash

. vars.cfg

docker commit ${CONTAINER_NAME} ${CONTAINER_NAME}-backup:$(date '+%Y%m%d')
tar cvzf ${CONTAINER_NAME}-root.tar.gz root/
aws s3 cp ${CONTAINER_NAME}-root.tar.gz s3://focusdockertest/${CONTAINER_NAME}/root.tar.gz
rm -f ${CONTAINER_NAME}-root.tar.gz
