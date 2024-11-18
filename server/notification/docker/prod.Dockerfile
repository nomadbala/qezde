FROM golang:1.23.2-alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o notification-service .

FROM alpine:3.20.3 AS prod

WORKDIR /app

COPY --from=builder /build/notification-service ./notification-service
COPY --from=builder /build/.env ./.env

RUN apk add --no-cache tini && \
    chmod +x ./notification-service

ENTRYPOINT ["/sbin/tini", "--", "./notification-service"]