# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go application
RUN go build -o crud-storage-api cmd/main.go

# Expose port 8080
EXPOSE 8080

# Run the app
CMD ["./crud-storage-api"]
