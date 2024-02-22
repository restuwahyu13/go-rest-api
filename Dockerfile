# Builder Stage
FROM golang:latest as builder

# Set environment variables
ENV USER=appuser

# Create a non-root user
RUN useradd -u 10001 ${USER}

# Set the working directory
WORKDIR /usr/src/app

# Copy the Go modules files
COPY go.mod go.sum ./

# Add missing entries to go.sum
RUN go get github.com/mattn/go-isatty@v0.0.12
RUN go get github.com/sendgrid/rest@v2.6.3+incompatible
RUN go get golang.org/x/sys/unix@latest
RUN go get golang.org/x/net/context@latest

# Update and tidy Go modules
RUN go mod tidy
RUN go mod download

# Copy the rest of the application code
COPY . ./

# Build the Go application
ENV GO111MODULE="on" \
    GOARCH="amd64" \
    GOOS="linux" \
    CGO_ENABLED="0"

RUN go build -o main .

# Final Stage
FROM alpine:latest

# Set environment variables
ENV USER=appuser

# Create a non-root user
RUN adduser -D -u 10001 ${USER}

# Set the working directory
WORKDIR /usr/src/app

# Install necessary packages
RUN apk --no-cache add ca-certificates

# Copy the built executable from the builder stage
COPY --from=builder /usr/src/app/main .

# Change ownership to the non-root user
RUN chown ${USER}:${USER} ./main

# Switch to the non-root user
USER ${USER}

# Expose the port
EXPOSE 4000

# Command to run the executable
CMD ["./main"]
