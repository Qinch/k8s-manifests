VERSION ?= v0.1.0
PROJECT_NAME=app-gateway
PROJECT_CONF_NAME=app-gateway-conf


export VERSION
export PROJECT_NAME
export PROJECT_CONF_NAME
export DOCKER_USER
export DOCKER_PWD

.PHONY: ci build image
ci: build image manifests 

build:   
	./build/build.sh
image:
ifndef DOCKER_USER
	$(error 错误：必须传入DOCKER_USER参数)
endif

ifndef DOCKER_PWD
	$(error 错误：必须传入DOCKER_PWD参数)
endif
	./build/image.sh
manifests:
	./build/kustomize.sh
