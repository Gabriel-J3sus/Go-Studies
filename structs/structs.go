package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

type messageToSend struct {
	message   string
	sender    sender
	recipient user
}

type sender struct {
	user      // extends user struct
	rateLimit int
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}

	if mToSend.recipient.name == "" {
		return false
	}

	if mToSend.sender.number == 0 {
		return false
	}

	if mToSend.recipient.number == 0 {
		return false
	}

	return true
}

// getBasicAuth has a receiver of (authI authenticationInfo)
func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		authI.username,
		authI.password,
	)
}

func test(mToSend messageToSend) {
	fmt.Printf("Sending: '%s' from: (%v) to %s (%v)\n", mToSend.message, mToSend.sender.number, mToSend.recipient.name, mToSend.recipient.number)
	fmt.Println("=======================================================")
}

func main() {
	canSendMessage(messageToSend{
		sender: sender{
			user: user{
				name:   "Gabriel",
				number: 122312312312,
			},
			rateLimit: 500,
		},
		recipient: user{
			name:   "Cabral",
			number: 3247349237498,
		},
		message: "What's up, nigga?",
	})

	authenticationInfo.getBasicAuth(authenticationInfo{
		username: "Gabriel",
		password: "123456",
	})

	test(messageToSend{
		sender: sender{
			user: user{
				name:   "Gabriel",
				number: 122312312312,
			},
			rateLimit: 500,
		},
		recipient: user{
			name:   "Cabral",
			number: 3247349237498,
		},
		message: "What's up, nigga?",
	})

	// Anonymous structs - use only if this is used once
	// myCar := struct {
	// 	Make  string
	// 	Model string
	// }{
	// 	Make:  "Tesla",
	// 	Model: "Model 3",
	// }
}
