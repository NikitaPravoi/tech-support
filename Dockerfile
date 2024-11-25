FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/server/main.go

FROM debian:bullseye-slim

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the config.yml OR config.json file into the container
# COPY --from=builder /app/config/config.yml ./config/config.yml

# Expose port
EXPOSE 8000

# Command to run the executable
CMD ["./main"]
