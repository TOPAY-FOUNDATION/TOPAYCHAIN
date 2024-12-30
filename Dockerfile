# Use an official Golang image as the base image
FROM golang:1.20 as builder

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o blockchain main.go

# Use a minimal base image for the runtime
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the built binary from the builder
COPY --from=builder /app/blockchain .

# Expose the application's default port
EXPOSE 8080

# Command to run the application
CMD ["./blockchain"]
