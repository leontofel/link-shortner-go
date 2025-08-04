# syntax=docker/dockerfile:1

FROM golang:1.24-alpine

WORKDIR /app

# Optional: Install git for go get if needed
RUN apk add --no-cache git

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o link-shortener ./cmd/server

EXPOSE 8080

CMD ["./link-shortener"]
