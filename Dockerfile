############################
# STEP 1: Build executable binary
############################
FROM golang:1.22.4-alpine AS builder

# Install necessary dependencies
RUN apk update && \
    apk add --no-cache git && \
    apk add --no-cache ca-certificates

# Set the working directory
WORKDIR /app

# Copy all source code into the container
COPY . .

# Build the Go application
RUN go build -o gh5-backend

############################
# STEP 2: Build a small image
############################
FROM scratch

# Copy necessary files and certificates from builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/gh5-backend /app/gh5-backend
COPY --from=builder /app/.env* /app/
COPY --from=builder /app/gh5-bucket-service-account.json /app/

# Set the working directory
WORKDIR /app

# Command to run the application
CMD ["./gh5-backend"]
