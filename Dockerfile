# syntax=docker/dockerfile:1

# ---------- Build stage ----------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git for fetching private modules if needed
RUN apk add --no-cache git

# Copy go.mod and go.sum first (for caching go mod download)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build optimized binary
RUN go build -o link-shortener -ldflags="-s -w" ./cmd/server

# ---------- Run stage ----------
FROM alpine:latest

WORKDIR /app
RUN apk --no-cache add ca-certificates

# Copy the binary from builder stage
COPY --from=builder /app/link-shortener .

EXPOSE 8080
CMD ["./link-shortener"]
