# Use the official Golang image as the base image
FROM golang:1.18-alpine


WORKDIR /app
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

RUN go build -o main .

EXPOSE 8080

# Start the application
CMD ["./main"]