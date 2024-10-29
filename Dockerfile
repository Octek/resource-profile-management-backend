# Start from golang base image
FROM golang:alpine as builder

# Install necessary dependencies
RUN apk update && apk add --no-cache git make gcc libtool musl-dev ca-certificates dumb-init build-base

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod tidy && go mod download

# Initialize and update submodules
RUN git submodule update --init --recursive

# Copy the source code
COPY . .

# Build the Go app
RUN GOOS=linux go build -o main .

# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app/

# Copy necessary files from the builder
COPY --from=builder /app/seed_data.json ./seed_data.json
COPY --from=builder /app/main .
# Expose port
EXPOSE 4001

# Command to run the executable
CMD ["./main"]
