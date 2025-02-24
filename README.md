Simple example to get started with Kafka in Golang. This sample is tested for macOS

Follow these steps to create a record and consume it.

From the root of the directory - 

1. Run `docker compose up -d` to run docker in detach mode. This will install zookeeper and kafka as docker containers on your machine locally.
2. Start the consumer.
   `go run ./cmd/consumer`
3. Start the producer to create a record.
   `go run ./cmd/producer`
