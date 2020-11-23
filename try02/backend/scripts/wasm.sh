#!/usr/bin/env bash
set -euox pipefail
SCRIPT_DIR=$(dirname "$0")
# shellcheck disable=SC2086
echo ${SCRIPT_DIR}
cd "${SCRIPT_DIR}" && cd ../wasm

GOOS=js GOARCH=wasm go build -o ../asset/go.wasm
#GOOS=js GOARCH=wasm tinygo build -o ../asset/go.wasm
