#!/bin/bash

. vars.cfg

# save.sh
#docker exec -it ${CONTAINER_NAME} /usr/bin/tar --same-owner -cvzf user_backup.tar.gz /home/
#docker cp ${CONTAINER_NAME}:/etc/gitlab-runner/config.toml root/etc/gitlab-runner/config.toml
#docker cp ${CONTAINER_NAME}:/user_backup.tar.gz root/
docker cp ${CONTAINER_NAME}:/root/.fropilr root/root
docker cp ${CONTAINER_NAME}:/root/.gnupg root/root
