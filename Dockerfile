ARG GO_VERSION

FROM golang:${GO_VERSION} as builder
WORKDIR /webhooks
COPY . ./
RUN make install
RUN make test
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /webhooks/build/chatstatz-webhooks .
ENTRYPOINT [ "./chatstatz-webhooks" ]
CMD [ "--help" ]
