package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Pls enter your first name: ")
	userLastName := getUserData("Pls enter your last name: ")
	userBirthdate := getUserData("Pls enter your birth date (MM/DD/YYYY): ")

	// var appUser user

	// long instantiating
	appUser := user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}
	outputUserDetails(&appUser)

	// short instantiating
	// appUser = user{
	// 	userFirstName,
	// 	userLastName,
	// 	userBirthdate,
	// 	time.Now(),
	// }

	// other way instantiating
	// appUser = user{}
	// appUser.firstName = userFirstName
	// appUser.lastName = userLastName
	// appUser.birthdate = userBirthdate
	// appUser.createdAt = time.Now()

	// fmt.Println(firstName, lastName, birthdate)
}

func outputUserDetails(u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
