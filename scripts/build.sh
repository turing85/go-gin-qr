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
  go test \
    -cover \
    -coverpkg=./... \
    -coverprofile=profile.cov \
    ./... \
    && go tool cover --func profile.cov
  echo "----------------------------------------"
  echo "Building application"
  CGO_ENABLED=0 `# Enable static linking` \
    go build \
      -ldflags="-s -w" `# omit symbol table (-s) and DWARF symbol table (-w)` \
      -o app \
      -tags netgo `# Necessary so that the executable is self-contained (i.e. fully statically linked)`
  echo "========================================"
}

function build() {
  cd ..
  build_go
  build_container "Containerfile"
}

build "${@}"