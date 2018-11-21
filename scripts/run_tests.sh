#!/usr/bin/env bash

set -e

printf "\n==> Installing dependencies\n"
make install

printf "\n==> Running tests...\n"
make test
