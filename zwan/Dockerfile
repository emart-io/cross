FROM golang:1.20 as builder

WORKDIR /go/src
COPY . .
RUN make build

FROM alpine
#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
#RUN apk --no-cache add ca-certificates
#COPY bundles/zwan /usr/local/bin
COPY --from=builder /go/src/bundles/zwan /usr/local/bin
ENTRYPOINT ["/usr/local/bin/zwan"]

EXPOSE 50051