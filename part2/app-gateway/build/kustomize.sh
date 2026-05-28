#!/bin/sh
set -e

BASE_PATH=$(cd "$(dirname "$0")" || exit; pwd)
OVERLAYS_PATH="${BASE_PATH}/../deployments/overlays"
MANIFESTS_PATH="${BASE_PATH}/../deployments/manifests"


SUBDIRS=$(find "${OVERLAYS_PATH}" -mindepth 1 -maxdepth 1 -type d)
for subdir in ${SUBDIRS}; do
    cd ${subdir} || exit
    kustomize edit set image mock-image="${DOCKER_USER}/${PROJECT_NAME}":"${VERSION}" 
    kustomize edit set image mock-conf-image="${DOCKER_USER}/${PROJECT_CONF_NAME}":"${VERSION}"
    cd -

    dir_name=$(basename "${subdir}")
    output_dir="${MANIFESTS_PATH}/${dir_name}"
    mkdir -p "${output_dir}"
    echo "kustomize build：${subdir}"
    kustomize build "${subdir}" -o "${output_dir}"
done
