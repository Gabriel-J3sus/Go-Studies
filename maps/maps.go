package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("Invalid sizes")
	}

	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name:       name,
			phoneNumer: phoneNumber,
		}
	}

	return userMap, nil
}

// Maps are deleted through a pointer, that's why I'm not returning it
func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	existingUser, ok := users[name]
	if !ok {
		return false, errors.New("Not found")
	}

	if existingUser.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

type user struct {
	name                 string
	phoneNumer           int
	scheduledForDeletion bool
}

func test(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("------------------------------")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value: \n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumer)
	}
}

func testDel(users map[string]user, name string) {
	defer fmt.Println("------------------------------")
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}

	if deleted {
		fmt.Println("Deleted user:", name)
	} else {
		fmt.Println("Did not delete user:", name)
	}

	println("================================")
	println("Final map keys:")
	for _, user := range users {
		fmt.Println(" - name:", user.name)
	}
}

// Maps work the same way as a javascript object
func main() {
	test([]string{"John", "Bob"}, []int{12345678, 123456})

	users := map[string]user{
		"Gabriel": {
			name:                 "Gabriel",
			phoneNumer:           12345678,
			scheduledForDeletion: true,
		},
		"Phil": {
			name:                 "Phil",
			phoneNumer:           12345678,
			scheduledForDeletion: false,
		},
	}

	testDel(users, "Phil")
	testDel(users, "Gabriel")
}
