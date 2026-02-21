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

func (u user) outputUserData() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearOutput() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) *user {
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
}

func main() {
	firstName := getUserData("Please enter first name: ")
	lastName := getUserData("Please enter last name: ")
	birthDate := getUserData("Please enter birthday (DD/MM/YYYY): ")

	var appUser *user

	appUser = newUser(firstName, lastName, birthDate)

	appUser.outputUserData()
	appUser.clearOutput()
	appUser.outputUserData()
}

func getUserData(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scan(&value)

	return value
}
