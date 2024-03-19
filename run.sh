#!/bin/sh

PROJECT_DIR="/go/src/api"

cd "$(dirname "$0")/" || exit 1

docker run -it --rm \
    --name di-api-local \
    -w "${PROJECT_DIR}" \
    -e "API_DIR=/go/src/api" \
    -e "air_wd=${PROJECT_DIR}" \
    -v "$(pwd)"/:"${PROJECT_DIR}" \
    -v "$GOPATH"/pkg/mod:/go/pkg/mod \
    --network=host \
    cosmtrek/air:v1.40.4
