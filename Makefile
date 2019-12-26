GO_VERSION=1.13.5
IMAGE_TAG=latest
IMAGE_NAME=chatstatz-webhooks
GCLOUD_CONTAINER_HOST=gcr.io
GCLOUD_PROJECT_ID=chatstatz-project

.DEFAULT_GOAL=.help
.SILENT: ;

install: ## Install dependencies
	CGO_ENABLED=0 go get ./...

build: ## Build chatstatz-webhooks server
	CGO_ENABLED=0 GOOS=linux go build -o chatstatz-webhooks .

test: ## Run tests
	CGO_ENABLED=0 go test -v

docker-build: ## Build chatstatz-webhooks image
	printf "==> Building $(IMAGE_NAME) image... "
	docker build --build-arg GO_VERSION=$(GO_VERSION) -t $(IMAGE_NAME):$(IMAGE_TAG) . >/dev/null
	printf "Done.\r\n"

docker-push: ## Publish chatstatz-webhooks images to repository
	printf "==> Tagging and pushing $(IMAGE_NAME) image... "
	docker tag $(IMAGE_NAME) $(GCLOUD_CONTAINER_HOST)/$(GCLOUD_PROJECT_ID)/$(IMAGE_NAME):$(IMAGE_TAG)
	docker push $(GCLOUD_CONTAINER_HOST)/$(GCLOUD_PROJECT_ID)/$(IMAGE_NAME):$(IMAGE_TAG) >/dev/null
	printf "Done.\r\n"
