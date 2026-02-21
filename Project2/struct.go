package main

import (
	"fmt"
	"struct-app/user"
)

func main() {
	firstName := getUserData("Please enter first name: ")
	lastName := getUserData("Please enter last name: ")
	birthDate := getUserData("Please enter birthday (DD/MM/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserData()
	appUser.ClearOutput()
	appUser.OutputUserData()
}

func getUserData(text string) string {
	fmt.Print(text)
	var value string
	fmt.Scanln(&value)

	return value
}
