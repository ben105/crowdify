package message_queue

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

const PartitionAny = kafka.PartitionAny

type Producer interface {
	Produce(b []byte)
}

type KafkaProducer struct {
	p     *kafka.Producer
	topic string
}

func NewProducer(broker, topic string) Producer {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": broker,
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to create producer: %v\n", err))
	}
	return &KafkaProducer{p: producer, topic: topic}
}

func (kp *KafkaProducer) Produce(b []byte) {
	defer kp.p.Close()

	deliveryChan := make(chan kafka.Event)

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &kp.topic,
			Partition: PartitionAny,
		},
		Value: b,
	}

	err := kp.p.Produce(message, deliveryChan)
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
