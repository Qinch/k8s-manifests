#!/bin/bash
set -euo pipefail

BASE_PATH=$(cd "$(dirname "$0")" || exit; pwd)
PROTO_PATH="${BASE_PATH}/../proto"
echo ${PROTO_PATH}
PROTO_SUBDIR=$(find ${PROTO_PATH} -type d -not -path "*google*"   -not -path "${PROTO_PATH}")
for proto in ${PROTO_SUBDIR}; do
  echo ${proto}
  protoc \
    --proto_path="${PROTO_PATH}" \
    --proto_path=/usr/local/include \
    --go_out="${PROTO_PATH}" \
    --go_opt=paths=source_relative \
    --go-grpc_out="${PROTO_PATH}" \
    --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out="${PROTO_PATH}" \
    --grpc-gateway_opt=paths=source_relative \
   ${proto}/*.proto
done

