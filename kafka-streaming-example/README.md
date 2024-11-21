# Kafka Streaming Example

## How It Works

    Input Topic: Messages are produced to input-topic (e.g., using the Kafka Console Producer).

    Streaming Application:
    * Reads messages from input-topic.
    * Processes messages (e.g., converts text to uppercase).
    * Produces transformed messages to output-topic.

    Output Topic: Transformed messages are consumed for validation or further processing.

## Key Features

    * Demonstrates real-time message processing.
    * Implements a basic Kafka streaming pipeline.
    * Showcases decoupling of services using Kafka topics.

## Next steps

    This example can be expanded with more complex transformations, integrating with databases, or chaining multiple services to build a robust streaming pipeline.

## How to run

1. Start the services

   ```bash
       docker-compose up --build

   ```

2. Create Kafka topics:

   ```bash
       docker exec -it kafka kafka-topics --create --topic input-topic --bootstrap-server kafka:9092 --partitions 1 --replication-factor 1
       docker exec -it kafka kafka-topics --create --topic output-topic --bootstrap-server kafka:9092 --partitions 1 --replication-factor 1
   ```

3. Produce test data to input-topic:

   ```bash
       docker exec -it kafka kafka-console-producer --topic input-topic --bootstrap-server kafka:9092
   ```

4. Type messages in the console

5. Consume transformed data from output-topic:

   ```bash
       docker exec -it kafka kafka-console-consumer --topic output-topic --bootstrap-server kafka:9092 --from-beginning
   ```

_Hint: Access Prometheus at <http://localhost:9090> and Grafana at <http://localhost:3000>_

6. Observe:

- Metrics in Prometheus (<http://localhost:9090/targets>).
- Visualizations in Grafana (Kafka dashboard).
