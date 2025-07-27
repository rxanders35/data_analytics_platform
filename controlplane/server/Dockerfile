FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . . 

RUN  CGO_ENABLED=0 go build -o  /app/controlplane

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/controlplane /app/controlplane

EXPOSE 8080

CMD ["/app/controlplane"]