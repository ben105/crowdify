package kafka

import (
	"fmt"
	"log"

	"github.com/ben105/crowdify/packages/env"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Publish(m string) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": env.Broker,
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %v\n", err)
	}
	defer producer.Close()

	deliveryChan := make(chan kafka.Event)

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &env.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(m),
	}

	err = producer.Produce(message, deliveryChan)
	if err != nil {
		log.Fatalf("Failed to produce message: %v\n", err)
		return
	}

	e := <-deliveryChan
	msg := e.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		log.Fatalf("Delivery failed: %v\n", msg.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)
	}
	close(deliveryChan)
}
