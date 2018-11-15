# Chatstatz Webhooks Service

The webhooks service listens for subscribed Twitch webhook events and
pushes event payloads to the message for consumption by the ACP server.

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
