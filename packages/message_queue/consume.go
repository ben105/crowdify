package message_queue

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer interface {
	Initialize() error
	ReadMessage() (*Message, error)
	CommitOffsets(offsets []TopicPartition) error
	Close() error
}

type KafkaConsumer struct {
	c     *kafka.Consumer
	topic string
}

type Subscription struct {
	Topic   string
	GroupId string
}

func NewConsumer(broker string, s Subscription) Consumer {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  broker,
		"group.id":           s.GroupId,
		"enable.auto.commit": true,
		"auto.offset.reset":  "earliest",
		"socket.timeout.ms":  10000,
		"session.timeout.ms": 60000,
	}
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		panic(fmt.Sprintf("Failed to create consumer: %v\n", err))
	}
	return &KafkaConsumer{c: consumer, topic: s.Topic}
}

func (k *KafkaConsumer) ReadMessage() (*Message, error) {
	msg, err := k.c.ReadMessage(-1)
	if err != nil {
		return nil, err
	}
	topicPartition := TopicPartition{
		Topic:     *msg.TopicPartition.Topic,
		Partition: msg.TopicPartition.Partition,
		Offset:    int64(msg.TopicPartition.Offset),
	}
	headers := make([]Header, len(msg.Headers))
	for i, header := range msg.Headers {
		headers[i] = Header{
			Key:   header.Key,
			Value: header.Value,
		}
	}
	return NewMessage(msg.Key, msg.Value, topicPartition, headers...), nil
}

func (k *KafkaConsumer) Initialize() error {
	return k.c.SubscribeTopics([]string{k.topic}, nil)
}

func (k *KafkaConsumer) Close() error {
	return k.c.Close()
}

func (k *KafkaConsumer) CommitOffsets(offsets []TopicPartition) error {
	tps := make([]kafka.TopicPartition, len(offsets))
	for i, offset := range offsets {
		tps[i] = kafka.TopicPartition{
			Topic:     &offset.Topic,
			Partition: offset.Partition,
			Offset:    kafka.Offset(offset.Offset + 1), // +1 because committed offset is the next expected one
		}
	}
	_, err := k.c.CommitOffsets(tps)
	return err
}
