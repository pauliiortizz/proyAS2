# Use the Go image with Alpine for building and running the application
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Dockerfile de search-api
COPY wait-for.sh /wait-for.sh
RUN chmod +x /wait-for.sh
CMD /wait-for.sh rabbitmq 5672 -- go run main.go

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the rest of the code and build the application
COPY . .
RUN go build -o app ./main.go

# Expose the port on which the app will run
EXPOSE 8081

# Command to run the binary
CMD ["./app"]