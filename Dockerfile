# Gorona Dockerfile

FROM golang:1.13-alpine3.11

RUN apk update && \
    apk add --no-cache git=2.24.1-r0 && \
    go get -u github.com/harik8/gorona

ENTRYPOINT ["gorona"]