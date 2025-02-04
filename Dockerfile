FROM alpine:latest

WORKDIR /app/test

RUN apk add --no-cache wget go \
    && wget -qO- https://taskfile.dev/install.sh | sh -s -- -b /usr/local/bin 

ENV GOCACHE=/app/.cache

RUN adduser -D  -h /app tester \
    && mkdir -p ${GOCACHE} \
    && chmod -R a+w ${GOCACHE}

USER tester

ENTRYPOINT ["task", "test"]