# Stage 1: Build the Go binary
FROM golang:1.16 AS builder

WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main

# Stage 2: Create a minimal runtime image
FROM scratch

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main /app/main

# Expose the port the application will listen on
EXPOSE 8000

# Define the command to run the application
CMD ["/app/main"]
