FROM alpine:latest

WORKDIR /app/test

RUN apk add task, go

ENTRYPOINT ["task", "test"]