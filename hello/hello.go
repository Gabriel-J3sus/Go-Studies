package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// message, err := greetings.Hello("Jorge")
	// message, err := greetings.Hello("")

	names := []string{
		"John",
		"Peter",
		"Mike",
	}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(message)
	fmt.Println(messages)
	for _, message := range messages {
		fmt.Println(message)
	}
}
