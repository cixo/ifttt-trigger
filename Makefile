.PHONY: default

DOCKER_IMG_NAME=ifttt-trigger

DOCKER=docker
COMMIT_ID=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date '+%Y%m%d')
DOCKER_IMG_VERSION=$(APP_VERSION)-$(COMMIT_ID)-$(BUILD_DATE)
DOCKER_IMG_FULL_NAME_TAG=$(DOCKER_IMG_NAME):$(DOCKER_IMG_VERSION)
DOCKER_IMG_FULL_NAME_LATEST=$(DOCKER_IMG_NAME):latest

build:
	$(DOCKER) build -t $(DOCKER_IMG_NAME) .
inspect:
	docker images | grep $(DOCKER_IMG_NAME)
	$(DOCKER) history $(DOCKER_IMG_NAME)
run:
	$(DOCKER) run --rm -it --name "$(DOCKER_IMG_NAME)_app" $(DOCKER_IMG_NAME) "event_123"
