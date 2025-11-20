# ---------- STAGE 1: Build ----------
FROM docker.io/library/golang:1.25-alpine AS builder

# Install git (dibutuhkan go mod download)
RUN apk add --no-cache git

# Set workdir
WORKDIR /app

# Copy go mod & sum
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary (disable CGO → binary portable)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server cmd/server/main.go


# ---------- STAGE 2: Runtime ----------
FROM alpine:3.19

# Create non-root user
RUN adduser -D appuser
USER appuser

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/server /app/server

# Copy config folder (untuk config.dev.yaml / config.uat.yaml / config.prod.yaml)
COPY configs/ /app/config/

# Expose port (sesuaikan Gin)
EXPOSE 8080

# Set default environment (bisa override saat run)
ENV APP_ENV=dev

# Start app
CMD ["./server"]
