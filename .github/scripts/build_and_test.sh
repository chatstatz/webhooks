#!/usr/bin/env bash

set -euo pipefail

docker build --pull --no-cache -t ${GCP_IMAGE_NAME} .
