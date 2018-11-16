# Chatstatz Webhooks Service

The webhooks service listens for subscribed Twitch webhook events and
pushes event payloads to the message for consumption by the ACP server.

[![Build status](https://badge.buildkite.com/03f303396fe05d51a5d7e420915544dde6a316fb8b32dc2012.svg)](https://buildkite.com/chatstatz/chatstatz-webhooks)

## Development

### Prerequisites

- Docker
- Go v1.11 or greater

### Make Commands

```txt
make                List commands
make install        Install dependencies
make test           Run tests
make build          Build chatstatz-webhooks server
make docker-build   Build chatstatz_webhooks image
make docker-push    Publish chatstatz_webhooks images to repository
```
