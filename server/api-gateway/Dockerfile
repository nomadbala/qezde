FROM golang:1.23.2-alpine AS builder

WORKDIR /build

COPY . /build

RUN go mod download

COPY . .

RUN go build -o api-gateway .

FROM alpine:3.20.3 AS hoster
COPY --from=builder /build/.env ./.env
COPY --from=builder /build/api-gateway ./api-gateway

ENTRYPOINT ["./api-gateway"]

EXPOSE ${APP_PORT}