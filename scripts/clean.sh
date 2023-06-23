#!/usr/bin/env bash
set -e

cd "$(dirname "$(realpath -s "$0")")" || exit 1

function clean() {
  echo "========================================"
  echo "Running go clean"
  go clean
  echo "----------------------------------------"
  echo "Removing auxiliary files"
  cd ..
  rm -rf app
  rm -rf app.*
  rm -rf profile.cov
  echo "========================================"
}

clean "${@}"