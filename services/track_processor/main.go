package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ben105/crowdify/packages/db"
	"github.com/ben105/crowdify/packages/kafka"
)

var conn *db.DbConnection

func main() {
	conn = db.Connect()

	consumer, err := kafka.NewKafkaConsumer(func(msg *kafka.Message, results chan<- kafka.Message) {
		if msg.Error == nil {
			var track *db.UnprocessedTrack
			err := json.Unmarshal(msg.KafkaMessage.Value, &track)
			if err != nil {
				fmt.Printf("Error unmarshalling message: %v\n", err)
				return
			}
			processTrack(track)
			// TODO: Add error handling.
			results <- *msg
		} else {
			fmt.Printf("Consumer error: %v (%v)\n",
				msg.Error,
				msg.KafkaMessage)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	consumer.Start(-1)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigchan
	fmt.Printf("Caught signal %v: terminating consumer...\n", sig)
	consumer.Stop()

	fmt.Println("Consumer stopped. Exiting...")
}
