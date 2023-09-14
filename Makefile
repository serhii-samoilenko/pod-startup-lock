
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

dep:
	@echo ">>> Make: Updating dependencies"
	dep ensure -v

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

docker-push-latest:
	@echo ">>> Make: Pushing all docker images with latest tag"
	$(MAKE) -C k8s-health docker-push-latest
	$(MAKE) -C init docker-push-latest
	$(MAKE) -C lock docker-push-latest
