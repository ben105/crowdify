package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ben105/crowdify/packages/db"
	"github.com/ben105/crowdify/packages/env"
	"github.com/ben105/crowdify/packages/message_queue"
)

type MessageProcessor struct{}

func (mp *MessageProcessor) ProcessMessage(msg *message_queue.Message, results chan<- message_queue.Result) {
	var track *db.UnprocessedTrack
	err := json.Unmarshal(msg.Value, &track)
	if err != nil {
		log.Fatalf("Error unmarshalling message: %v\n", err)
		return
	}
	processTrack(track)
	// TODO: Add error handling.
	results <- message_queue.Result{
		Message: msg,
	}
}

var conn *db.DbConnection

func main() {
	conn = db.Connect()

	topic := env.GetTopic()
	if topic == "" {
		log.Fatal("Kafka topic is not set")
	}

	processor := new(MessageProcessor)
	consumer := message_queue.NewConsumer(env.GetBroker(), message_queue.Subscription{Topic: env.GetTopic(), GroupId: env.GetGroupId()})
	runner := message_queue.NewRunner(consumer, processor, *message_queue.NewCommitManager(5 * time.Second))

	runner.Start()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigchan
	fmt.Printf("Caught signal %v: terminating consumer...\n", sig)
	runner.Stop()

	fmt.Println("Consumer stopped. Exiting...")
}
