package main

import (
	"fmt"
	"strings"
)

func removeProfanity(message *string) {
	messageVal := *message // dereference message value
	messageVal = strings.ReplaceAll(messageVal, "merda", "***")
	messageVal = strings.ReplaceAll(messageVal, "cacete", "***")
	messageVal = strings.ReplaceAll(messageVal, "caralho", "***")

	*message = messageVal // update address with new value
}

func test(messages []string) {
	for _, message := range messages {
		removeProfanity(&message)
		fmt.Println(message)
	}
}

func main() {
	test([]string{
		"Que merda!",
		"To com sede, cacete!",
		"Vem aqui, caralho!",
	})
}
