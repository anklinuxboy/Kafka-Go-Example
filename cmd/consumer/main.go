package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	config := kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "my-group",
	}

	consumer, err := kafka.NewConsumer(&config)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer consumer.Close()

	topic := "my-topic"
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	fmt.Printf("listening for messages for topic %s\n", topic)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("received message %s %s", string(*msg.TopicPartition.Topic), string(msg.Value))
		} else {
			fmt.Printf("error :%s", err.Error())
		}
	}
}
