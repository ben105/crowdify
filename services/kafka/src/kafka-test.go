package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args

	if args[1] == "publish" {
		if len(args) < 3 {
			log.Fatal("Expected a message to publish.")
			os.Exit(1)
		}
		publish(args[2])
	} else if args[1] == "consume" {
		consume()
	} else {
		log.Fatalf("Expected required argument [publish|consume] but got %s", args[1])
	}
}
