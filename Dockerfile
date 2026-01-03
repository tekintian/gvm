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

# Generate version file from git (if tag exists), otherwise create default version
# 如果 app_build/version.go 已存在(由 CI/CD 预先生成),则直接使用;否则从 git 生成
RUN if [ -f app_build/version.go ]; then \
      echo "Using pre-generated version file"; \
    else \
      echo "Generating version from git..."; \
      go run app_build/gen_version.go || \
      (mkdir -p app_build && \
       echo 'package app_build' > app_build/version.go && \
       echo '' >> app_build/version.go && \
       echo 'const (' >> app_build/version.go && \
       echo '  // ShortVersion 短版本号（Docker 构建）' >> app_build/version.go && \
       echo '  ShortVersion = "0.0.0-docker"' >> app_build/version.go && \
       echo ')' >> app_build/version.go); \
    fi

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
