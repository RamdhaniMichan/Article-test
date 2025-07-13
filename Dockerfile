# Build stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install required build dependencies
RUN apk add --no-cache git

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application with optimization flags
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main cmd/main.go

# Final stage
FROM alpine:3.18

WORKDIR /app

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates tzdata

# Copy the binary from builder
COPY --from=builder /app/main .
COPY --from=builder /app/migration ./migration
COPY .env.example .env

# Set the timezone
ENV TZ=Asia/Jakarta

# Expose port
EXPOSE 8081

# Run the binary
CMD ["./main"]