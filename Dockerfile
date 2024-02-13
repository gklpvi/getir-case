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

## Build the Go app, if you want to
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

ENTRYPOINT ["go", "run", "main.go"]
