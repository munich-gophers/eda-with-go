
FROM golang:1.23-alpine

WORKDIR /app
COPY main.go .
RUN go mod init consumer && go mod tidy && go build -o consumer .

CMD ["./consumer"]
