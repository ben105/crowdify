package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ben105/crowdify/services/kafka"
)

func main() {
	args := os.Args

	if args[1] == "publish" {
		if len(args) < 3 {
			log.Fatal("Expected a message to publish.")
			os.Exit(1)
		}
		kafka.Publish([]byte(args[2]))
	} else if args[1] == "consume" {
		consumer, err := kafka.NewKafkaConsumer()
		if err != nil {
			log.Fatal(err)
		}
		consumer.Start(-1)
		for msg := range consumer.Messages {
			if msg.Error == nil {
				fmt.Printf("Message on %s: %s\n",
					msg.KafkaMessage.TopicPartition,
					string(msg.KafkaMessage.Value))
			} else {
				fmt.Printf("Consumer error: %v (%v)\n",
					msg.Error,
					msg.KafkaMessage)
			}
		}
	} else {
		log.Fatalf("Expected required argument [publish|consume] but got %s", args[1])
	}
}
