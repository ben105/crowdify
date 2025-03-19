package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ben105/crowdify/packages/db"
	"github.com/ben105/crowdify/services/kafka"
)

var conn *db.DbConnection

func main() {
	conn = db.Connect()

	consumer, err := kafka.NewKafkaConsumer()
	if err != nil {
		log.Fatal(err)
	}

	consumer.Start(-1)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigchan
		fmt.Printf("Caught signal %v: terminating consumer...\n", sig)
		consumer.Stop()
	}()

	fmt.Println("Listening for messages...")
	for msg := range consumer.Messages {
		if msg.Error == nil {
			fmt.Printf("Message on %s: %s\n",
				msg.KafkaMessage.TopicPartition,
				string(msg.KafkaMessage.Value))
			var track *db.UnprocessedTrack
			err := json.Unmarshal(msg.KafkaMessage.Value, &track)
			if err != nil {
				fmt.Printf("Error unmarshalling message: %v\n", err)
				continue
			}
			processTrack(track)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n",
				msg.Error,
				msg.KafkaMessage)
		}
	}
	fmt.Println("Consumer stopped. Exiting...")
}
