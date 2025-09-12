FROM golang:1.19 AS builder

# Set work directory
WORKDIR /backend

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app (will fail if no main.go, but this is a shell setup)
RUN CGO_ENABLED=0 GOOS=linux go build -o backend

# Final minimal image
FROM alpine:latest
WORKDIR /backend
COPY --from=builder /backend/backend ./backend

# Expose API port
EXPOSE 8000

# Run the binary (will fail if not built, but this is a shell setup)
CMD ["./backend"]
