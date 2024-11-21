# How it works

## Acknowledgements

    * When a message is processed successfully, the consumer explicitly acknowledges the message using msg.Ack(false).
    * If processing fails, the message is rejected using msg.Nack(false, true), and RabbitMQ requeues it for retry.

## Simulated Error

    If the message body contains the string "error", the consumer logs an error and requeues the message.

## Benefits

    Highlights manual message acknowledgment and error handling.
    Demonstrates RabbitMQ's capability to requeue unacknowledged messages.

## How to

1. Run

```bash
    docker-compose up --build.
```

2. Publish messages using the producer:

```bash
    docker-compose exec producer ./producer
```

3. Watch the consumer-ack logs to observe message acknowledgment and retry behavior.
