# Versions
GO_VERSION=1.11.2

# General
GCLOUD_PROJECT=chatstatz-project
BUILD_TAG=latest

# Docker
DOCKER_HOST=asia.gcr.io
DOCKER_IMAGE=chatstatz_webhooks

.DEFAULT_GOAL=.help
.SILENT: ;

install: ## Install dependencies
	CGO_ENABLED=0 go get

build: ## Build chatstatz-webhooks server
	CGO_ENABLED=0 GOOS=linux go build -o chatstatz-webhooks .

docker-build: ## Build chatstatz_webhooks image
	printf "==> Building $(DOCKER_IMAGE) image... "
	docker build --build-arg GO_VERSION=$(GO_VERSION) -t $(DOCKER_IMAGE):$(BUILD_TAG) . >/dev/null
	printf "Done.\r\n"

docker-push: ## Publish chatstatz_webhooks images to repository
	printf "==> Tagging and pushing $(DOCKER_IMAGE) image... "
	docker tag $(DOCKER_IMAGE) $(DOCKER_HOST)/$(GCLOUD_PROJECT)/$(DOCKER_IMAGE):$(BUILD_TAG)
	docker push $(DOCKER_HOST)/$(GCLOUD_PROJECT)/$(DOCKER_IMAGE):$(BUILD_TAG) >/dev/null
	printf "Done.\r\n"

.help:
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'
