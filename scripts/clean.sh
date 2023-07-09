#!/usr/bin/env bash
set -e

cd "$(dirname "$(realpath -s "$0")")" || exit 1
source build-commons.sh

cd ..
clean "${@}"