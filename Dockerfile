# Build stage
FROM golang:1.22.2-alpine AS builder
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o keyvalue-cache .

# Final stage
FROM alpine:latest
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/keyvalue-cache .

# Expose the required port
EXPOSE 7171

# Run the binary
CMD ["./keyvalue-cache"]
