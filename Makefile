.SILENT:;
.PHONY: default install test build

default:
	echo "TODO..."

install:
	CGO_ENABLED=0 go get ./...

build:
	rm -rf build
	CGO_ENABLED=0 GOOS=linux go build -o build/chatstatz-webhooks ./cmd/webhooks

test:
	CGO_ENABLED=0 go test -v ./...
