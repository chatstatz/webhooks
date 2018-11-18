# Chatstatz Webhooks Service

The webhooks service listens for subscribed Twitch webhook events and
pushes event payloads to the message for consumption by the ACP server.

[![Build status](https://badge.buildkite.com/03f303396fe05d51a5d7e420915544dde6a316fb8b32dc2012.svg)](https://buildkite.com/chatstatz/chatstatz-webhooks)

## Development

### Prerequisites

- Docker
- Google Cloud CLI tool
- Go v1.11 or greater

### Setup Commands

```bash
# Login in to your GCP account
glcoud auth login

# Change project to "chatstatz-project" if not already configured
gcloud config set project "chatstatz-project"

# Configure Docker to use gcloud as a credentials helper
gcloud auth configure-docker
```

### Make Commands

```txt
make                            List commands
make install                    Install dependencies
make test                       Run tests
make build                      Build chatstatz-webhooks server
make docker-build               Build chatstatz_webhooks image
make docker-push                Publish chatstatz_webhooks images to repository
gcloud-delete-untagged-images   Delete untagged chatstatz_webhooks GCR images
```
