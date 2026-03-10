package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	//Connect to broker
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	message := []byte("Hello World - Pub")

	//Publish message to subject
	err = nc.Publish("hello", message)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message published!")
}
