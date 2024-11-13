# Stage 1: Builder
FROM golang:1.20 as builder

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod tidy
RUN go mod download

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
