FROM alpine:latest

WORKDIR /app/test

RUN apk add --no-cache wget tar \
    && GO_VERSION=$(wget -qO- https://go.dev/VERSION?m=text | head -n 1) \
    && wget -q https://go.dev/dl/${GO_VERSION}.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf ${GO_VERSION}.linux-amd64.tar.gz \
    && rm ${GO_VERSION}.linux-amd64.tar.gz \
    && wget -qO- https://taskfile.dev/install.sh | sh -s -- -b /usr/local/bin 

ENV PATH=$PATH:/usr/local/go/bin \
    GOCACHE=/app/.cache

RUN adduser -D  -h /app tester \
    && mkdir -p  ${GOCACHE} \
    && chmod -R a+w ${GOCACHE}

USER tester


ENTRYPOINT ["task", "test"]
