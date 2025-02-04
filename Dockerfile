FROM alpine:latest

WORKDIR /app/test

RUN apk add --no-cache wget go \
    && wget -qO- https://taskfile.dev/install.sh | sh -s -- -b /usr/local/bin 

ENTRYPOINT ["task", "test"]