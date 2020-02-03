#!/usr/bin/env bash

set -euo pipefail

SLICE=$(git rev-parse --abbrev-ref HEAD)
SLICE_VERSION=$(git rev-parse --short HEAD)

if [ "${SLICE}" == "master" ]; then
    docker tag ${GCP_IMAGE_NAME} ${GCP_CONTAINER_HOST}/${GCP_PROJECT_ID}/${GCP_IMAGE_NAME}:latest
    docker push ${GCP_CONTAINER_HOST}/${GCP_PROJECT_ID}/${GCP_IMAGE_NAME}:latest
else
    docker tag ${GCP_IMAGE_NAME} ${GCP_CONTAINER_HOST}/${GCP_PROJECT_ID}/${GCP_IMAGE_NAME}:${SLICE}-${SLICE_VERSION}
    docker push ${GCP_CONTAINER_HOST}/${GCP_PROJECT_ID}/${GCP_IMAGE_NAME}:${SLICE}-${SLICE_VERSION}
fi
