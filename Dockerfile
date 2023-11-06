FROM alpine:3.17
RUN apk add ca-certificates bash libc6-compat && update-ca-certificates && rm /var/cache/apk/*
COPY ./bin/manager /usr/local/bin/manager

ENTRYPOINT ["/usr/local/bin/manager"]
