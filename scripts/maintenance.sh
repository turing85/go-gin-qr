#!/usr/bin/env bash

cd "$(dirname "$(realpath -s "$0")")" || exit 1

function maintenance() {
  cd ..
  echo "========================================"
  echo "Running go get"
  go get \
    -t `# get test dependencies` \
    -u `# update modules` \
    ./...
  echo "----------------------------------------"
  echo "Running go vet"
  go vet ./...
  echo "----------------------------------------"
  echo "Running go tidy"
  go mod tidy
  echo "----------------------------------------"
  echo "Running go fmt"
  go fmt ./...
  echo "========================================"
}

maintenance