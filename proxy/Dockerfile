FROM golang:1.21.4-alpine AS builder

WORKDIR /app
COPY go.mod proxy.go ./

RUN go build -o proxy

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/proxy .
CMD ["./proxy"]
