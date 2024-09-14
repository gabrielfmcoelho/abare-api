# Use a lightweight base image
FROM golang:1.20-alpine AS builder

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory inside the container
WORKDIR /app

# Install necessary dependencies
RUN apk --no-cache add gcc g++ make sqlite sqlite-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the application
RUN go build -o main .

# Use a minimal image for the final stage
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Install SQLite on the minimal image
RUN apk --no-cache add sqlite

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8000
EXPOSE 8000

# Command to run the application
CMD ["./main"]
