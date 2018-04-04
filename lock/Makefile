# Go build params
GOOS ?= linux
export GOOS
export CGO_ENABLED=0

# App info for docker build
app ?= app.info
include $(app)
export $(shell sed 's/=.*//' $(app))

# Default docker url
DOCKER_URL ?= local

.DEFAULT_GOAL := build

test:
	@echo ">>> Make $(APP_NAME): Running tests"
	go test -cover -v ./...

build: test
	@echo ">>> Make $(APP_NAME): Building"
	go build -a -o bin/$(APP_NAME)

docker-build: build
	@echo ">>> Make $(APP_NAME): Building docker image"
	docker build -t "$(DOCKER_URL)/$(DOCKER_NAME):$(APP_VERSION)" .

docker-push: docker-build
	@echo ">>> Make $(APP_NAME): Pushing docker image"
	docker login -u $(DOCKER_USER) -p $(DOCKER_PW) $(DOCKER_URL)
	docker push "$(DOCKER_URL)/$(DOCKER_NAME):$(APP_VERSION)"
