GO_VERSION=1.11.2
IMAGE_TAG=latest
IMAGE_NAME=chatstatz_webhooks
GCLOUD_HOST=asia.gcr.io
GCLOUD_PROJECT=chatstatz-project

.DEFAULT_GOAL=.help
.SILENT: ;

install: ## Install dependencies
	CGO_ENABLED=0 go get

build: ## Build chatstatz-webhooks server
	CGO_ENABLED=0 GOOS=linux go build -o chatstatz-webhooks .

test: ## Run tests
	CGO_ENABLED=0 go test -v

docker-build: ## Build chatstatz_webhooks image
	printf "==> Building $(IMAGE_NAME) image... "
	docker build --build-arg GO_VERSION=$(GO_VERSION) -t $(IMAGE_NAME):$(IMAGE_TAG) . >/dev/null
	printf "Done.\r\n"

docker-push: ## Publish chatstatz_webhooks images to repository
	printf "==> Tagging and pushing $(IMAGE_NAME) image... "
	docker tag $(IMAGE_NAME) $(GCLOUD_HOST)/$(GCLOUD_PROJECT)/$(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(GCLOUD_HOST)/$(GCLOUD_PROJECT)/$(IMAGE_NAME):$(IMAGE_TAG) >/dev/null
	printf "Done.\r\n"

.help:
	grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-10s\033[0m %s\n", $$1, $$2}'
