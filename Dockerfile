FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./

COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /monitoringo

# Run
CMD ["/monitoringo"]