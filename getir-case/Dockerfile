# Use the official Golang image as base
FROM golang:latest AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Build a minimal distribution image
FROM alpine:latest

# Copy the binary from the builder stage
COPY --from=builder /app/app /app/app

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["/app/app"]
