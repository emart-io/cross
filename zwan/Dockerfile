FROM alpine
#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN apk --no-cache add ca-certificates
COPY bundles/zwan /usr/local/bin
#COPY --from=0 /go/emart/zbay /usr/local/bin
ENTRYPOINT ["/usr/local/bin/zwan"]

EXPOSE 50051