FROM ubuntu:latest
LABEL authors="Bima Dewantoro"

ENTRYPOINT ["top", "-b"]

# Use the official Golang image as the base image
FROM golang AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download and cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Copy the .env file
COPY .env .env

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build image to the final container
COPY --from=build /app/main .

# Expose the port that your application listens on
EXPOSE 8080

# Start your application with environment variables from .env file
CMD ["sh", "-c", "source .env && ./main"]
