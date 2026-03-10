package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to broker
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	defer nc.Close()

	// Subscribe to subject
	_, err = nc.Subscribe("hello", func(msg *nats.Msg) {
		log.Println("Received message:", string(msg.Data))
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Waiting for messages...")

	select {} // ?? - seems to keep things running
}
