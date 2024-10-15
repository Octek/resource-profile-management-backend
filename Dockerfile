# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
RUN apk add build-base
RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
#COPY go.mod go.sum ./

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod tidy
RUN go mod download

# Update the submodule for the app so latest changes can be applied
RUN git submodule update

# Build the Go app
RUN GOOS=linux go build -o main .

# Start a new stage from scratch (means creating a new docker and copying the required information from the above mentioned one)
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/seed_data.json ./seed_data.json
COPY --from=builder /app/locales/en-US.json ./locales/en-US.json
COPY --from=builder /app/main .
#COPY --from=builder /app/.env .

# Expose port 3106 to the outside world
EXPOSE 4001

#Command to run the executable
CMD ["./main"]