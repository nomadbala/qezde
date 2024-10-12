FROM golang:1.23.1-alpine as builder

WORKDIR /build

COPY . /build

RUN go mod download

COPY . .

RUN go build -o auth-service .

FROM alpine:3.18 as hoster
COPY --from=builder /build/.env ./.env
COPY --from=builder /build/auth-service ./auth-service

ENTRYPOINT ["./auth-service"]

EXPOSE 8081