# Stage 1: Build the Go application
FROM golang:1.22 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app with static linking
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main cmd/main.go

# Stage 2: Run the Go application
FROM alpine:latest

# Install necessary CA certificates for Alpine
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/main .

COPY .env .

COPY templates ./templates

# Ensure the binary is executable
RUN chmod +x ./main

# Expose port 8002 to the outside world
EXPOSE 8002

# Command to run the executable
CMD ["./main"]
