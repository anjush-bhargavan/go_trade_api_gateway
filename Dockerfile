# Stage 1: Build the Go application
FROM golang:1-alpine AS builder

LABEL maintainer="Anjush Bhargavan <anjushbhargavan12@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Cross-compile for ARM64 architecture
RUN GOARCH=arm64 GOOS=linux go build -o api_gateway

# Stage 2: Run the Go application using a scratch image
FROM alpine:latest

LABEL maintainer="Anjush Bhargavan <anjushbhargavan12@gmail.com>"

# Install CA certificates package to enable TLS certificate verification
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/api_gateway .
COPY --from=builder /app/templates ./templates
COPY .env .

EXPOSE 8080

CMD ["./api_gateway"]

