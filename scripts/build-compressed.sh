#!/usr/bin/env bash
set -e

cd "$(dirname "$(realpath -s "$0")")" || exit 1
source build-commons.sh

function build_go() {
  echo "========================================"
  echo "Getting dependencies"
  go clean
  go get \
    -d `# only get dependencies, do not install them` \
    -t `# get test dependencies` \
    -v `# print the name of the packages`
  echo "----------------------------------------"
  echo "Running tests"
  go test
  echo "----------------------------------------"
  echo "Building application"
  go build \
    -ldflags="-s -w" `# omit symbol table (-s) and DWARF symbol table (-w)` \
    -o app \
    -tags netgo
  echo "----------------------------------------"
  echo "Compressing application"
  upx --brute app
  echo "========================================"
}

function build() {
  cd ..
  BUILD_CONTAINER="yes"
  build_go
  build_container "Containerfile"
}

build "${@}"