package main

import (
	"errors"
)

func getMessageWithRetries() [3]string {
	// array - fixed size
	return [3]string{
		"click here to sign up",
		"pretty, please click here",
		"we beg you to sign up",
	}
}

// slice - dynamic array, most used, better developer experience but less performance
// pointer behaviour
const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()
	if plan == planPro {
		//slicing to get all itens and turn into an slice
		return allMessages[:], nil
	}

	if plan == planFree {
		return allMessages[0:2], nil
	}

	return nil, errors.New("Unsupported plan")
}

func getMessagesCosts(messages []string) []float64 {
	costs := make([]float64, len(messages)) // creates an empty array
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01
		costs[i] = cost
	}

	return costs
}

func main() {
	getMessageWithRetries()
	getMessageWithRetriesForPlan(planFree)
}
