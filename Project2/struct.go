package main

import (
	"fmt"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
}

func main() {
	firstName := getUserData("Please enter first name: ")
	lastName := getUserData("Please enter last name: ")
	birthDate := getUserData("Please enter birthday (DD/MM/YYYY): ")

	outputUserData(user.firstName, user.lastName, user.birthDate)

}

func outputUserData(firstName, lastName, birthDate user) {
	fmt.Println(firstName, lastName, birthDate)
}

func getUserData(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scan(&value)

	return value
}
