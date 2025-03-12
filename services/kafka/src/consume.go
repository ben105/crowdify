package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func consume() {
	// Create a new Kafka consumer configuration
	config := &kafka.ConfigMap{
		"bootstrap.servers": broker,
		"group.id":          groupId,
		"group.instance.id": groupInstanceId,
		"auto.offset.reset": "earliest",
	}

	// Create a new consumer
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	defer consumer.Close()

	// Subscribe to the topic
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	fmt.Println("Kafka Consumer started... Press Ctrl+C to exit.")

	// Capture system interrupts to gracefully shut down the consumer
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigchan
		fmt.Printf("Caught signal %v: terminating consumer...\n", sig)
		os.Exit(1)
	}()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
