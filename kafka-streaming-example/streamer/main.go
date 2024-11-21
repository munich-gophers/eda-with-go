package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func main() {
	// Kafka Reader to consume from "input-topic"
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"kafka:9092"},
		Topic:    "input-topic",
		GroupID:  "streamer-group",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer reader.Close()

	// Kafka Writer to produce to "output-topic"
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"},
		Topic:   "output-topic",
	})
	defer writer.Close()

	log.Println("Streaming application started...")

	startPrometheusServer()
	cleanup := initTracing()
	defer cleanup()

	log.Println("Tracing started...")

	tracer := otel.Tracer("streamer")
	_, span := tracer.Start(context.Background(), "ProcessMessage")
	defer span.End()

	for {
		start := time.Now()

		// Read and process messages
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Encountered an error", err)
			errorMessages.Inc()
			span.RecordError(err)
			continue
		}
		processedMessages.Inc()

		// Simulate processing logic
		span.AddEvent("Message processed", trace.WithAttributes(
			attribute.String("message.key", string(msg.Key)),
			attribute.String("message.value", string(msg.Value)),
		))
		processedValue := strings.ToUpper(string(msg.Value))
		time.Sleep(100 * time.Millisecond) // Simulate processing time
		messageProcessingTime.Observe(time.Since(start).Seconds())

		// Write processed message
		writer.WriteMessages(context.Background(), kafka.Message{
			Key:   msg.Key,
			Value: []byte(processedValue),
		})
	}
}
