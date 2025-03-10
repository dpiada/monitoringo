FROM golang:1.24.1

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy all source files
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o monitoringo

# Run the built binary
CMD ["/app/monitoringo"]
