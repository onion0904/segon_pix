# syntax=docker/dockerfile:1

FROM golang

# Set destination for COPY
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project and wait-for-it.sh script
COPY . .

# Build the Go application
RUN go build -o app

# Command to wait for DB and then run the application
CMD ["./app"]

