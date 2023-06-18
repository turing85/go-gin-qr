#!/usr/bin/env bash

cd "$(dirname "$(realpath -s "$0")")" || exit 1

function maintenance() {
  cd ..
  go get -u
  go vet
  go mod tidy
}

maintenance