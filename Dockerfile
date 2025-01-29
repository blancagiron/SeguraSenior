
FROM alpine:latest

WORKDIR /app/test

RUN apk update && apk add --no-cache \
    curl \
    wget \
    tar \
    bash \
    && GO_VERSION=$(curl -sL https://go.dev/VERSION?m=text | head -n 1) \
    && wget https://go.dev/dl/${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf ${GO_VERSION}.linux-amd64.tar.gz \
    && rm ${GO_VERSION}.linux-amd64.tar.gz \
    && curl -sL https://taskfile.dev/install.sh | sh -s -- -b /usr/local/bin \
    && apk del curl wget tar

ENV PATH=$PATH:/usr/local/go/bin \
    GOCACHE=/tmp/go-cache \
    GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN adduser -D -h /app tester \
    && mkdir -p /app/.cache \
    && chown -R tester:tester /app/.cache

USER tester


ENTRYPOINT ["task", "test"]
