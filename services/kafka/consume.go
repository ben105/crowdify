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
	Messages chan Message
	quit     chan struct{}
	wg       sync.WaitGroup
	once     sync.Once
}

func NewKafkaConsumer() (*KafkaConsumer, error) {
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
		consumer: consumer,
		Messages: make(chan Message),
		quit:     make(chan struct{}),
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
		for {
			fmt.Println("Polling for messages...")
			select {
			case <-kc.quit:
				fmt.Println("Stopping consumer...")
				kc.once.Do(func() { close(kc.Messages) })
				return
			default:
				msg, err := kc.consumer.ReadMessage(pollTimeout)
				fmt.Println("Message received")
				if err != nil {
					if kErr, ok := err.(k.Error); ok && !kErr.IsTimeout() {
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
	kc.wg.Wait()
	kc.once.Do(func() { close(kc.Messages) })
	return kc.consumer.Close()
}
