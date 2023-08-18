FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN make build

FROM alpine:3.10
COPY --from=builder /app/bundles/zwan /usr/local/bin
ENTRYPOINT ["/usr/local/bin/zwan"]

EXPOSE 50051