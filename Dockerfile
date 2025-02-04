FROM alpine:latest

RUN apk update && apk add --no-cache wget go \
    && wget -qO- https://taskfile.dev/install.sh | sh -s -- -b /usr/local/bin 

ENV GOCACHE=/app/.cache

RUN adduser -D  -h /app tester \
    && mkdir -p ${GOCACHE} \
    && chmod -R a+w ${GOCACHE}

WORKDIR /app/test

USER tester

ENTRYPOINT ["task", "test"]