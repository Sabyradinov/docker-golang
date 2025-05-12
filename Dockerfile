# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod tidy

COPY . .

RUN go build -o /app/server .

# Stage 2: Set up Nginx to serve as a reverse proxy
FROM nginx:alpine

# Copy the built Go server binary
COPY --from=builder /app/server /usr/bin/server

# Expose the port Nginx will listen on
EXPOSE 80

# Start both Nginx and the Go server
CMD ["/bin/sh", "-c", "/usr/bin/server & nginx -g 'daemon off;'"]