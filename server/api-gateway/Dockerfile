FROM golang:1.23.1-alpine AS builder

WORKDIR /build

COPY . /build

RUN go mod download

COPY . .

RUN go build -o api-gateway .

FROM alpine:3.18 as hoster
COPY --from=builder /build/.env ./.env
COPY --from=builder /build/api-gateway ./api-gateway

ENTRYPOINT ["./api-gateway"]

EXPOSE 8080