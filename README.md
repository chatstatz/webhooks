# webhooks

The webhooks service listens for subscribed Twitch webhook events and
pushes event payloads to the NATS message queue.

## Development

### Prerequisites

- Docker
- Go v1.13 or greater

### CI Pipeline

For continuous integration (CI) this project uses [Cloud Build](https://cloud.google.com/cloud-build/).
The build pipeline will run tests and publish Docker images to [GCR](https://cloud.google.com/container-registry/).

### Make Commands

```txt
make install                    Install dependencies
make test                       Run tests
make build                      Build chatstatz-webhooks server
```
