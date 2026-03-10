package main

import "fmt"

func main() {
	sendsSoFar := 430
	const sendsToAdd = 25
	sendsSoFar = incrementalSends(sendsSoFar, sendsToAdd)
	fmt.Println("You've sent", sendsSoFar, "messages")

	// ----------------------------------

	firstName, _ := getNames() // completly ignore the param so the compiler does not throw unnused variable errors
	// firstName, lastName := getNames()
	fmt.Println("Welcome", firstName)
}

func incrementalSends(sendsSoFar, sendsToAdd int) int {
	sendsSoFar += sendsToAdd
	return sendsSoFar
}

func getNames() (string, string) {
	return "Manuel", "Gomes"
}

// ----------------------------------
// *** Named return values ***
// use wisely, only on small functions
// func getCoords() (int, int) {
// 	var x int
// 	var y int
// 	return x, y
// }

// // is the same as:
// func getCoords2() (x, y int) {
// 	// x and y are initialized with zero values
// 	return //automatically returns x and y
// }

// ----------------------------------
