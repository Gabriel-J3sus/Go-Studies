package main

import (
	"fmt"
	"time"
)

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

type message interface {
	getMessage() string
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime)
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v`, sr.reportName)
}

func test(m message) {
	sendMessage(m)
	fmt.Println("===============================")
}

// Interfaces are collections of method signatures.
// A type "implements" an interface if it has all of the methods of the given interface defined on it
func main() {
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})

	test(birthdayMessage{
		recipientName: "Jesus",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
}
