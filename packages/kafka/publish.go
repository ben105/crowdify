package kafka

import (
	"fmt"
	"log"

	"github.com/ben105/crowdify/packages/env"
	k "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Publish(m []byte) {
	producer, err := k.NewProducer(&k.ConfigMap{
		"bootstrap.servers": env.GetBroker(),
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %v\n", err)
	}
	defer producer.Close()

	deliveryChan := make(chan k.Event)

	topic := env.GetTopic()
	message := &k.Message{
		TopicPartition: k.TopicPartition{Topic: &topic, Partition: k.PartitionAny},
		Value:          m,
	}

	err = producer.Produce(message, deliveryChan)
	if err != nil {
		log.Fatalf("Failed to produce message: %v\n", err)
		return
	}

	e := <-deliveryChan
	msg := e.(*k.Message)
	if msg.TopicPartition.Error != nil {
		log.Fatalf("Delivery failed: %v\n", msg.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset)
	}
	close(deliveryChan)
}
