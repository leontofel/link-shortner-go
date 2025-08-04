# syntax=docker/dockerfile:1

# ---------- Build stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git for go get (needed for private repos or some modules)
RUN apk add --no-cache git

# Cache go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build with optimizations for smaller binary
RUN go build -o link-shortener -ldflags="-s -w" ./cmd/server

# ---------- Run stage ----------
FROM alpine:latest

WORKDIR /app

# Install CA certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Copy only the binary from the builder stage
COPY --from=builder /app/link-shortener .

EXPOSE 8080

CMD ["./link-shortener"]
