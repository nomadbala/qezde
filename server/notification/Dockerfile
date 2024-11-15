FROM golang:1.23.2-alpine AS builder

WORKDIR /build

COPY . /build

RUN go mod download

COPY . .

RUN go build -o notification-service .

FROM alpine:3.18 AS hoster
COPY --from=builder /build/.env ./.env
COPY --from=builder /build/notification-service ./notification-service

ENTRYPOINT ["./notification-service"]

EXPOSE 8082