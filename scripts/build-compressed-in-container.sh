#!/usr/bin/env bash
set -e
cd "$(dirname "$(realpath -s "$0")")" || exit 1

source build-commons.sh

function build() {
  cd ..
  BUILD_CONTAINER="yes"
  build_container "Containerfile.build-and-compress-in-container"
}

build