FROM golang:latest

WORKDIR /app

# Copy the Go modules files
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

COPY . .

RUN go build -o server ./cmd/server

EXPOSE 50052

CMD ["./server"]
