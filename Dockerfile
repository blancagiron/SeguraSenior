FROM alpine:latest

WORKDIR /app/test

RUN apk add --no-cache wget go \
    && wget -qO- https://taskfile.dev/install.sh | sh -s -- -b /usr/local/bin 

ENV GOCACHE=/app/.cache

RUN mkdir -p ${GOCACHE} \
    && chmod -R a+w ${GOCACHE}

ENTRYPOINT ["task", "test"]