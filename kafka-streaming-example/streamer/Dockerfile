
FROM golang:1.23-alpine

WORKDIR /app
COPY . .
RUN go mod init streamer && go mod tidy && go build -o streamer .

CMD ["./streamer"]
