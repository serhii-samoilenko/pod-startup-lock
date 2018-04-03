# App info for docker build
app ?= version.info
include $(app)
export $(shell sed 's/=.*//' $(app))

.DEFAULT_GOAL := build

linux:
	@echo ">>> Make: Setting GO target OS to 'linux'"
	GOOS=linux
    export GOOS

darwin:
	@echo ">>> Make: Setting GO target OS to 'darwin'"
	GOOS=darwin
    export GOOS

windows:
	@echo ">>> Make: Setting GO target OS to 'windows'"
	GOOS=windows
    export GOOS

glide:
	@echo ">>> Make: Updating dependencies"
	glide update --strip-vendor --skip-test

build:
	@echo ">>> Make: Building all modules"
	$(MAKE) -C k8s-health
	$(MAKE) -C init
	$(MAKE) -C lock

docker-build:
	@echo ">>> Make: Building all docker images"
	$(MAKE) -C k8s-health docker-build
	$(MAKE) -C init docker-build
	$(MAKE) -C lock docker-build

docker-push:
	@echo ">>> Make: Pushing all docker images"
	$(MAKE) -C k8s-health docker-push
	$(MAKE) -C init docker-push
	$(MAKE) -C lock docker-push
