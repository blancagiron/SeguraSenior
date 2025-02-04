FROM alpine:latest

WORKDIR /app/test

RUN apk add --no-cache task go 

ENTRYPOINT ["task", "test"]