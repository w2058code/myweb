# --- Stage 1: Build the Go binary ---
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy dependency files first to leverage Docker layer caching
COPY go.mod ./
# COPY go.sum ./ (Uncomment if you have dependencies)
RUN go mod download

# Copy source code and build the production binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o web-server main.go

# --- Stage 2: Final minimal runtime ---
FROM alpine:3.22

# Add certificates for applications making external HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-compiled binary from the builder stage
COPY --from=builder /app/web-server .

# Vercel overrides the runtime port via the PORT environment variable
EXPOSE 80

# Execute the application
CMD ["./web-server"]
