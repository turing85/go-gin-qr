#!/usr/bin/env bash
set -e
cd "$(dirname "$(realpath -s "$0")")" || exit 1

source build-commons.sh

function build() {
  cd ..
  build_container "runner-from-builder"
}

build