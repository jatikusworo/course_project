# --------------------
# 1. Builder Stage
# --------------------
FROM docker.io/library/golang:1.25 AS builder

WORKDIR /app

# Copy dependency list
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

ARG BUILD_TIME
# Build static binary
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-buildid=${BUILD_TIME}" -o server cmd/server/main.go


# --------------------
# 2. Runtime Stage
# --------------------
FROM alpine:latest
LABEL build_time="${BUILD_TIME}"

WORKDIR /app

# Install timezone (optional)
RUN apk add --no-cache tzdata

# Copy binary
COPY --from=builder /app/server ./app

# Copy config folder (yaml)
COPY configs ./configs

EXPOSE 8080

# Use entrypoint for flexibility
ENTRYPOINT ["./app"]
