FROM golang:1.20-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ADD . .

RUN go build -o main ./cmd

FROM alpine:latest

RUN apk add --no-cache bash

COPY --from=builder /app/main /app/main
COPY --from=builder /app/.env /app/.env

EXPOSE 8080

CMD ["/app/main"]
