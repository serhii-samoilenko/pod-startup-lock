# Go build params
GOOS ?= linux
export GOOS
export CGO_ENABLED=0

# App info for docker build
app = app.info
include $(app)
export $(shell sed 's/=.*//' $(app))

# App version for docker build
ver = ../version.info
include $(ver)
export $(shell sed 's/=.*//' $(ver))

.DEFAULT_GOAL := build

test:
	@echo ">>> Make $(APP_NAME): Running tests"
	go test -cover -v ./...

build: test
	@echo ">>> Make $(APP_NAME): Building"
	go build -a -o bin/$(APP_NAME)

docker-build: build
	@echo ">>> Make $(APP_NAME): Building docker image"
	docker build -t $(DOCKER_USER)/$(DOCKER_NAME):$(APP_VERSION) .

docker-push: docker-build
	@echo ">>> Make $(APP_NAME): Pushing docker image"
	docker login -u $(DOCKER_USER) -p $(DOCKER_PW) $(DOCKER_URL)
	docker push $(DOCKER_USER)/$(DOCKER_NAME):$(APP_VERSION)

docker-push-latest: docker-push
	@echo ">>> Make $(APP_NAME): Tagging docker image version $(APP_VERSION) as latest and pushing"
	docker tag $(DOCKER_USER)/$(DOCKER_NAME):$(APP_VERSION) $(DOCKER_USER)/$(DOCKER_NAME):latest
	docker push $(DOCKER_USER)/$(DOCKER_NAME):latest
