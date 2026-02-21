package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter first name: ")
	lastName := getUserData("Please enter last name: ")
	birthDate := getUserData("Please enter birthday (DD/MM/YYYY): ")

	var appUser user

	appUser = user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}

	outputUserData(&appUser)
}

func outputUserData(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scan(&value)

	return value
}
