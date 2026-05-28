#!/bin/bash
set -e
# 1. 校验必填环境变量是否存在
if [[ -z "${PROJECT_NAME}" || -z "${VERSION}" || -z "${DOCKER_USER}" ]]; then
    echo "错误：请设置环境变量 PROJECT_NAME、VERSION、DOCKER_USER"
    exit 1
fi
echo "${DOCKER_USER}/${PROJECT_NAME}:${VERSION}"
docker build -t "${DOCKER_USER}/${PROJECT_NAME}:${VERSION}" --build-arg PROJECT_NAME=${PROJECT_NAME} .
