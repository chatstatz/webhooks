# webhooks

The webhooks service listens for subscribed Twitch webhook events and
pushes event payloads to the NATS message queue.

## Development

### Prerequisites

- Docker
- Google Cloud CLI tool
- Go v1.13 or greater

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
```
