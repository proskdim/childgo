FROM golang:1.23-alpine AS builder

RUN apk add --update \
    sqlite \
    libsqlite3-dev \
    gcc \
    musl-dev \
    rm -rf /var/cache/apk/*

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=1 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o childgo .

CMD [ "./childgo" ]