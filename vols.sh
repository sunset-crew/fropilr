directory="$(pwd)/example"
docker volume create --driver local \
    --opt type=none \
    --opt device=${directory} \
    --opt o=bind example
