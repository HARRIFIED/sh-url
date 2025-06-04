# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install necessary packages
RUN apk add --no-cache git

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Debug: Check what was copied and verify structure
RUN echo "=== Checking project structure ===" && \
    ls -lah && \
    echo "=== Checking cmd directory ===" && \
    ls -lah cmd/ && \
    echo "=== Checking cmd/server directory ===" && \
    ls -lah cmd/server/

# Build the application with CGO disabled for Alpine compatibility
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o url-shortener-go ./cmd/server

# Verify the binary was created
RUN ls -lah url-shortener-go

# Stage 2: Create a minimal image
FROM alpine:latest

# Install ca-certificates for HTTPS requests (if needed)
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/url-shortener-go .

# Copy the templates directory from the builder stage
COPY --from=builder /app/templates ./templates

# Make sure the binary is executable
RUN chmod +x ./url-shortener-go

# Verify the binary and templates are in the final image
RUN ls -lah ./url-shortener-go && ls -lah ./templates/

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./url-shortener-go"]