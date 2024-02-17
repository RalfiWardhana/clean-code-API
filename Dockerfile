# Use an official Golang runtime as a parent image
FROM golang:1.20 as builder

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Install required packages
RUN go mod tidy
RUN go mod download
RUN go mod verify

# Build the application
RUN go build -o main .

# Use an official Alpine Linux image as the final stage
FROM alpine:3.14

# Create the app directory
RUN mkdir -p /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main /app/

# Expose the port that the application is listening on
EXPOSE 8080

# Start the application
CMD ["/app/main"]
