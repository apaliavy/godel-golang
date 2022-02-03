FROM golang:1.17

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/gateway/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./main"]