
FROM golang:1.23-alpine

WORKDIR /app
COPY main.go .
RUN go mod init producer && go mod tidy && go build -o producer .

CMD ["./producer"]
