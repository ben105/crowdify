package main

import (
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
		kafka.Publish(args[2])
	} else if args[1] == "consume" {
		kafka.Consume()
	} else {
		log.Fatalf("Expected required argument [publish|consume] but got %s", args[1])
	}
}
