# Use an official Golang runtime as a base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

COPY . .

RUN go mod init my-go-app
RUN go get github.com/go-sql-driver/mysql
RUN go mod download

# Build the Go application
RUN go build -o main .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
