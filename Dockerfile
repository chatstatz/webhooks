ARG GO_VERSION

FROM golang:${GO_VERSION} as builder
WORKDIR /build
COPY . ./
RUN make install
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /chatstatz-webhooks
COPY --from=builder /build/chatstatz-webhooks .
ENTRYPOINT [ "./chatstatz-webhooks" ]
CMD [ "--help" ]
