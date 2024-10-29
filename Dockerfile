# Stage 1: Builder
FROM golang:alpine as builder

# Install necessary dependencies
RUN apk update && apk add --no-cache git make gcc libtool musl-dev ca-certificates dumb-init build-base

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod files (adjusted to avoid go.sum issue if it's missing)
COPY go.mod ./
# Copy go.sum if it exists
COPY go.sum ./
# Download all dependencies. `go mod tidy` will create go.sum if it's missing.
RUN go mod tidy && go mod download

# Initialize and update submodules (optional if submodules aren’t accessible in Docker)
# RUN git submodule update --init --recursive

# Copy the source code
COPY . .

# Build the Go app
RUN GOOS=linux go build -o main .

# Stage 2: Final Image
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
