package main

import (
	"errors"
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

func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name and birthdate are required!")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	firstName := getUserData("Please enter first name: ")
	lastName := getUserData("Please enter last name: ")
	birthDate := getUserData("Please enter birthday (DD/MM/YYYY): ")

	var appUser *user

	appUser, err := newUser(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserData()
	appUser.clearOutput()
	appUser.outputUserData()
}

func getUserData(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scanln(&value)

	return value
}
