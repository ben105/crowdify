package main

import (
	"log"
	"os"
	"time"

	"github.com/ben105/crowdify/packages/env"
	"github.com/ben105/crowdify/packages/message_queue"
)

type MessagePrinter struct{}

func (mp *MessagePrinter) ProcessMessage(msg *message_queue.Message, results chan<- message_queue.Result) {
	results <- message_queue.Result{
		Message: msg,
		Error:   nil,
	}
}

func main() {
	args := os.Args

	if args[1] == "produce" {
		if len(args) < 3 {
			log.Fatal("Expected a message to produce.")
			os.Exit(1)
		}
		p := message_queue.NewProducer(env.GetBroker(), env.GetTopic())
		p.Produce([]byte(args[2]))
	} else if args[1] == "consume" {
		c := message_queue.NewConsumer(env.GetBroker(), message_queue.Subscription{Topic: env.GetTopic(), GroupId: env.GetGroupId()})
		runner := message_queue.NewRunner(c, new(MessagePrinter), *message_queue.NewCommitManager(5 * time.Second))
		runner.Start()
		select {}
	} else {
		log.Fatalf("Expected required argument [publish|consume] but got %s", args[1])
	}
}
