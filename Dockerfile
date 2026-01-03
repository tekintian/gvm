# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git make

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -v -ldflags="-s -w" -o /app/gvm .

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS
RUN apk --no-cache add ca-certificates git

WORKDIR /root/.gvm

# Copy binary from builder
COPY --from=builder /app/gvm /usr/local/bin/gvm

# Make the binary executable
RUN chmod +x /usr/local/bin/gvm

# Set up environment
ENV GOROOT=/root/.gvm/go
ENV PATH=/root/.gvm/go/bin:${PATH}
ENV GVM_MIRROR=https://golang.google.cn/dl/

# Set entrypoint
ENTRYPOINT ["/usr/local/bin/gvm"]
CMD ["--help"]
