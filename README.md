# webhooks

The webhooks service listens for subscribed Twitch webhook events and
pushes event payloads to the NATS message queue.

## Environment Variables

This application uses environment variables for configurations.
Below are all expected variables and their default values.

| Variable | Default | Description |
|:---------|:-------:|:------------|
| `WEBHOOKS_HOST` | `127.0.0.1` | The host that the  webhooks server should run on |
| `WEBHOOKS_PORT` | `8080` | The port that the webhooks server should be served on |
| `NATS_HOST` | `0.0.0.0` | The NATS host address for NATS clients to connect to |
| `NATS_PORT` | `4222` | The NATS port for NATS clients to connect on |
| `NATS_QUEUE` | `twitch_channels` | The NATS queue for which to publish messages to |
| `LOG_LEVEL` | `info` | The log level to start logging from (see [here](https://github.com/chatstatz/logger)) |

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

## License

This repository is distributed under the terms of the [MIT](LICENSE) License.
