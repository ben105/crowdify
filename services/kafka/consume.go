package kafka

import (
	"fmt"
	"log"
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
	Messages chan Message
	quit     chan struct{}
}

func NewKafkaConsumer() (*KafkaConsumer, error) {
	config := &k.ConfigMap{
		"bootstrap.servers": env.GetBroker(),
		"group.id":          env.GetGroupId(),
		"group.instance.id": env.GetGroupInstanceId(),
		"auto.offset.reset": "earliest",
		// "socket.timeout.ms":  10000,
		// "session.timeout.ms": 60000,
	}
	consumer, err := k.NewConsumer(config)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer: consumer,
		Messages: make(chan Message),
		quit:     make(chan struct{}),
	}, nil
}

func (kc *KafkaConsumer) Start(pollTimeout time.Duration) {
	err := kc.consumer.SubscribeTopics([]string{env.GetTopic()}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Subscribed to topic %s\n", env.GetTopic())
	go func() {
		for {
			fmt.Println("Polling for messages...")
			select {
			case <-kc.quit:
				fmt.Println("Stopping consumer...")
				close(kc.Messages)
				return
			default:
				msg, err := kc.consumer.ReadMessage(pollTimeout)
				fmt.Println("Message received")
				if err != nil {
					if !err.(k.Error).IsTimeout() {
						kc.Messages <- Message{KafkaMessage: msg, Error: err}
					}
				} else {
					kc.Messages <- Message{KafkaMessage: msg, Error: nil}
				}
			}
		}
	}()
}

func (kc *KafkaConsumer) Stop() error {
	close(kc.quit)
	return kc.consumer.Close()
}
