VERSION ?= v0.1.0
PROJECT_NAME=app
DOCKER_USER=qinchaowhut

export VERSION
export PROJECT_NAME
export DOCKER_USER

.PHONY: ci build image
ci: build image 

build:   
	./build/build.sh
image:
	./build/image.sh
