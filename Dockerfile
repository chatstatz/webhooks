ARG GO_VERSION=1.13

FROM golang:${GO_VERSION} as builder
WORKDIR /app
COPY . ./
RUN make install
RUN make test
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
WORKDIR /app
COPY --from=builder /app/build/chatstatz-webhooks .
ENTRYPOINT [ "./chatstatz-webhooks" ]
