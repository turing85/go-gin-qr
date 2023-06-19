#!/usr/bin/env bash
set -e

function get_registry() {
  echo "${REGISTRY:-localhost}"
}

function get_registry_repository() {
  echo "${REGISTRY_REPO:-turing85}"
}

function get_image_name() {
  echo "${IMAGE_NAME:-go-gin-qr}"
}

function get_image_tag() {
  echo "${IMAGE_TAG:-latest}"
}

function has_command() {
  if ! [ -x "$(command -v "${1}")" ]
  then
    return 1
  fi
}

function has_buildah() {
  has_command buildah
}

function has_podman() {
  has_command podman
}

function has_docker() {
  has_command docker
}

function get_command() {
    if has_buildah
    then
      echo "buildah"
    elif has_podman
    then
      echo "podman"
    elif has_docker
    then
      echo "docker"
    else
      echo "Neither buildah, nor podman, nor docker has been found. Aborting."
      return 1
    fi
}

function construct_full_build_command() {
  local backslash=\\
  cat <<EOF
${1} build ${backslash}
  --file $(pwd)/build/${2} ${backslash}
  --format oci ${backslash}
  --tag "$(get_registry)/$(get_registry_repository)/$(get_image_name):$(get_image_tag)" ${backslash}
  --target runner ${backslash}
  $(pwd)
EOF
}

function should_build_container() {
  if [ -n "${BUILD_CONTAINER}" ]
  then
    return 0;
  else
    return 1;
  fi
}

function build_container() {
  if should_build_container
  then
    echo "========================================"
    echo "Building container"
    echo "----------------------------------------"
    local command;
    command=$(get_command)
    local full_build_command
    full_build_command=$(construct_full_build_command "${command}" "${1}")
    echo "found ${command}, starting build"
    echo "Full build command:"
    echo
    echo "$full_build_command"
    echo "========================================"
    eval "$full_build_command"
  fi
}