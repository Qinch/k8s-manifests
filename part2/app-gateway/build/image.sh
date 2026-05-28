#!/bin/sh
set -e

echo "${DOCKER_USER}/${PROJECT_NAME}:${VERSION}"
BASE_PATH=$(cd "$(dirname "$0")" || exit; pwd)
cd "${BASE_PATH}/.." || exit
docker build -t "${DOCKER_USER}/${PROJECT_NAME}:${VERSION}" --build-arg PROJECT_NAME=${PROJECT_NAME} .
docker login -u${DOCKER_USER} -p${DOCKER_PWD}
docker push "${DOCKER_USER}/${PROJECT_NAME}:${VERSION}"

cd "${BASE_PATH}/../configs" || exit
docker build -t "${DOCKER_USER}/${PROJECT_CONF_NAME}:${VERSION}" --build-arg PROJECT_NAME=${PROJECT_CONF_NAME} .
docker push "${DOCKER_USER}/${PROJECT_CONF_NAME}:${VERSION}"
