package main

import "fmt"

func main() {
	// initialize variables - empty initialization
	var smsSendingLimit int // same as smsSendingLimit := 0
	var costPerSMS float64  // same as costPerSMS := 0.0
	var hasPermission bool  // same as hasPermission := false
	var username string     // same as username := ""

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// declare multiple inline
	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"

	fmt.Print(averageOpenRate, displayMessage)

	// bool

	// string

	// byte // alias for uint8

	// rune // alias for int32
	// represents a Unicode code point

	// int int8 int16 int32 int64 // whole numbers
	// uint uint8 uint16 uint32 uint64 uintptr // positive whole numbers
	// float32 float64 // decimal numbers
	// complex64 complex128 // imaginary numbers (rare)

	// Parsing numbers
	temperatureInt := 88
	temperatureFloat := float64(temperatureInt) // 88.0

	accountAge := 2.6
	accountAgeInt := int(accountAge) // 2

	fmt.Println(temperatureFloat, accountAgeInt)

	// constants
	// - cannot use short declaration syntax
	const premiumPlanName = "Premium Plan"
	// premiumPlanName = "Basic Plan" // Throw error
	const basicPlanName = "Basic Plan" // works

	// - constants must be known at compile time
	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = minutesInHour * secondsInMinute
	fmt.Println("NUmber of seconds in an hour:", secondsInHour)

	// If statement initial state
	// if INITIAL_STATEMENT;  CONDITION {}
	// if length := getLength(email); length < 1 {
	// 	// length can be used inside this scope
	// 	fmt.Println("Email is invalid")
	// }
	//cannot use "length"

}
