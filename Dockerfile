FROM alpine:latest

WORKDIR /app

# Install CA certs for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the prebuilt binary from context
COPY link-shortener .

EXPOSE 8080

CMD ["./link-shortener"]
