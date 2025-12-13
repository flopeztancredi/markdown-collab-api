# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/bin/api ./cmd/api

# Runtime stage
FROM alpine:3.19

WORKDIR /app

# Add non-root user for security
RUN adduser -D -g '' appuser

# Copy binary from builder
COPY --from=builder /app/bin/api .

# Use non-root user
USER appuser

EXPOSE 8080

CMD ["./api"]
