.SILENT:;
.PHONY: install test build

RELEASE_VERSION:=$(shell git describe --always --long --dirty)

install:
	CGO_ENABLED=0 go get ./...

build:
	rm -rf build/ && \
	CGO_ENABLED=0 GOOS=linux go build -i -v -ldflags "-w -s -X main.version=${RELEASE_VERSION}" -o build/chatstatz-webhooks ./cmd/webhooks

test:
	CGO_ENABLED=0 go test -covermode=atomic ./...
