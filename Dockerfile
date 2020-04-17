# Gorona Dockerfile
FROM golang:1.13-alpine3.11

LABEL HARI KARTHIGASU <hari.karthigasu@gmail.com>

WORKDIR /app

COPY ./go.mod ./go.sum ./
COPY . .

RUN go build -o /usr/local/bin

ENTRYPOINT ["gorona"]