package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "demo-topic",
	})
	defer writer.Close()

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte("Key-1"),
		Value: []byte("Hello, Kafka from Producer!"),
	})
	if err != nil {
		log.Fatalf("Failed to write message: %v", err)
	}
	log.Println("Message sent successfully!")
}
