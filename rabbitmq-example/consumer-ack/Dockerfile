
FROM golang:1.23-alpine

WORKDIR /app
COPY main.go .
RUN go mod init consumer-ack && go mod tidy && go build -o consumer-ack .

CMD ["./consumer-ack"]
