#!/bin/bash
set -e
echo ${PROJECT_NAME}
echo ${VERSION}
BASE_PATH=$(cd "$(dirname "$0")" || exit; pwd)
cd "${BASE_PATH}/.."
go build  -o ./target/bin/${PROJECT_NAME} ./cmd/main.go
chmod +x ./target/bin/${PROJECT_NAME}


