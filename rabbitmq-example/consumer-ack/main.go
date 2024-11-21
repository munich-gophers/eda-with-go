package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare the queue (must match the producer)
	q, err := ch.QueueDeclare(
		"demo-ack-queue", // name
		true,             // durable
		false,            // delete when unused
		false,            // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack (set to false for manual acknowledgment)
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")

	// Process messages
	for msg := range msgs {
		log.Printf("Received a message: %s", msg.Body)

		// Simulate message processing
		if err := processMessage(msg.Body); err != nil {
			log.Printf("Failed to process message: %v", err)
			// Reject and requeue the message
			msg.Nack(false, true)
		} else {
			// Acknowledge the message
			msg.Ack(false)
			log.Println("Message acknowledged")
		}
	}
}

// Simulate message processing logic
func processMessage(body []byte) error {
	// Add custom processing logic here
	// Return an error if processing fails
	log.Printf("Processing message: %s", body)
	// Succes after 2 retries
	return nil
	// Error for retries
	// return errors.New("faulty processing")
}
