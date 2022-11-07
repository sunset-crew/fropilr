#!/bin/bash

. vars.cfg

# restore.sh
aws s3 cp s3://focusdockertest/${CONTAINER_NAME}/root.tar.gz .
tar xvzf root.tar.gz
rm -f root.tar.gz
