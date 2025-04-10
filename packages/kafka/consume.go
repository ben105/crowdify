package kafka

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ben105/crowdify/packages/env"
	k "github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Message struct {
	KafkaMessage *k.Message
	Error        error
}

type KafkaConsumer struct {
	consumer *k.Consumer

	processMessage func(msg *Message, results chan<- Message)

	results chan Message
	quit    chan struct{}

	wg sync.WaitGroup
}

func NewKafkaConsumer(processMessage func(msg *Message, results chan<- Message)) (*KafkaConsumer, error) {
	config := &k.ConfigMap{
		"bootstrap.servers":  env.GetBroker(),
		"group.id":           env.GetGroupId(),
		"enable.auto.commit": true,
		"auto.offset.reset":  "earliest",
		"socket.timeout.ms":  10000,
		"session.timeout.ms": 60000,
	}
	consumer, err := k.NewConsumer(config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer:       consumer,
		processMessage: processMessage,
		results:        make(chan Message),
		quit:           make(chan struct{}),
	}, nil
}

func (kc *KafkaConsumer) Start(pollTimeout time.Duration) {
	topic := env.GetTopic()
	if topic == "" {
		log.Fatal("Kafka topic is not set")
	}
	err := kc.consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Subscribed to topic %s on broker %s\n", topic, env.GetBroker())

	kc.wg.Add(1)
	go func() {
		defer kc.wg.Done()
		kc.processMessages(pollTimeout)
	}()

	kc.wg.Add(1)
	go func() {
		defer kc.wg.Done()
		kc.handleResults()
	}()

	fmt.Println("Kafka consumer started")
}

func (kc *KafkaConsumer) processMessages(pollTimeout time.Duration) {
	for {
		fmt.Println("Polling for messages...")
		select {
		case <-kc.quit:
			fmt.Println("Stopping consumer...")
			return
		default:
			msg, err := kc.consumer.ReadMessage(pollTimeout)
			fmt.Println("Message received")
			if err != nil {
				if kErr, ok := err.(k.Error); ok && !kErr.IsTimeout() {
					go kc.processMessage(&Message{KafkaMessage: msg, Error: err}, kc.results)
				}
			} else {
				go kc.processMessage(&Message{KafkaMessage: msg, Error: err}, kc.results)
			}
		}
	}
}

type TopicPartition struct {
	Topic     string
	Partition int32
}

func (kc *KafkaConsumer) handleResults() {
	offsetsToCommit := make(map[TopicPartition]k.Offset)
	commitTicker := time.NewTicker(5 * time.Second)
	defer commitTicker.Stop()
	for {
		select {
		case <-kc.quit:
			fmt.Println("Stopping result handler...")
			commitOffsets(kc.consumer, offsetsToCommit)
			return
		case <-commitTicker.C:
			fmt.Println("Commit ticker ticked")
			if len(offsetsToCommit) > 0 {
				fmt.Printf("Committing %d offsets\n", len(offsetsToCommit))
				commitOffsets(kc.consumer, offsetsToCommit)
				offsetsToCommit = make(map[TopicPartition]k.Offset)
			}
		case msg := <-kc.results:
			tp := TopicPartition{
				Topic:     *msg.KafkaMessage.TopicPartition.Topic,
				Partition: msg.KafkaMessage.TopicPartition.Partition,
			}
			if msg.Error != nil {
				fmt.Printf("Error processing message: %v\n", msg.Error)
				sendToDlq(msg.KafkaMessage, msg.Error)
			} else {
				currentOffset, exists := offsetsToCommit[tp]
				if !exists || msg.KafkaMessage.TopicPartition.Offset > currentOffset {
					offsetsToCommit[tp] = msg.KafkaMessage.TopicPartition.Offset
				}
				fmt.Printf("Processed message: %s\n", string(msg.KafkaMessage.Value))
			}
		}
	}
}

func commitOffsets(consumer *k.Consumer, offsetsToCommit map[TopicPartition]k.Offset) {
	if len(offsetsToCommit) == 0 {
		return
	}

	tps := make([]k.TopicPartition, 0, len(offsetsToCommit))

	for tp, offset := range offsetsToCommit {
		tps = append(tps, k.TopicPartition{
			Topic:     &tp.Topic,
			Partition: tp.Partition,
			Offset:    offset + 1, // +1 because committed offset is the next expected one
		})
	}

	log.Printf("Committing offsets for %d partitions", len(tps))

	_, err := consumer.CommitOffsets(tps)
	if err != nil {
		log.Printf("Failed to commit offsets: %v", err)
	} else {
		log.Printf("Successfully committed offsets")
	}
}

func (kc *KafkaConsumer) Stop() error {
	close(kc.quit)
	kc.wg.Wait()
	return kc.consumer.Close()
}

func sendToDlq(msg *k.Message, processingErr error) {
	producer, err := k.NewProducer(&k.ConfigMap{
		"bootstrap.servers": env.GetBroker(),
	})
	if err != nil {
		log.Printf("Failed to create DLQ producer: %v", err)
		return
	}
	defer producer.Close()
	headers := []k.Header{
		{
			Key:   "error_message",
			Value: []byte(processingErr.Error()),
		},
		{
			Key:   "original_topic",
			Value: []byte(*msg.TopicPartition.Topic),
		},
		{
			Key:   "original_partition",
			Value: []byte(fmt.Sprintf("%d", msg.TopicPartition.Partition)),
		},
		{
			Key:   "original_offset",
			Value: []byte(fmt.Sprintf("%d", msg.TopicPartition.Offset)),
		},
	}

	// Combine original headers with error headers
	if msg.Headers != nil {
		headers = append(headers, msg.Headers...)
	}

	// Send to DLQ topic
	dlqTopic := env.GetDeadLetterQueueTopic()
	deliveryChan := make(chan k.Event)

	err = producer.Produce(&k.Message{
		TopicPartition: k.TopicPartition{
			Topic:     &dlqTopic,
			Partition: k.PartitionAny,
		},
		Value:   msg.Value,
		Key:     msg.Key,
		Headers: headers,
	}, deliveryChan)

	if err != nil {
		log.Printf("Failed to produce DLQ message: %v", err)
		return
	}

	// Wait for delivery report
	e := <-deliveryChan
	m := e.(*k.Message)

	if m.TopicPartition.Error != nil {
		log.Printf("Failed to deliver message to DLQ: %v", m.TopicPartition.Error)
	} else {
		log.Printf("Message delivered to DLQ topic %s [%d] at offset %v",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
}
