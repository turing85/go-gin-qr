#!/usr/bin/env bash
set -e

cd "$(dirname "$(realpath -s "$0")")" || exit 1

function clean() {
  cd ..
  go clean
  rm -rf app
}

clean "${@}"