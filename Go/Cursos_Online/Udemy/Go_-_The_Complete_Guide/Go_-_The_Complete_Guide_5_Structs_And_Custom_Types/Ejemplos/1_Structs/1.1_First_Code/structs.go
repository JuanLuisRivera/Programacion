package main

import (
	"fmt"
	"time" //This package will allow us to use the time structand also some time related functions
)

type user struct { //We define a custom struct type "user" to store all the data from the user input
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time //Struct type provided by the "time" package
}

func main() {
	userFirstName := getUserData("Enter first name: ")
	userLastName := getUserData("Enter last name: ")
	userBirthdate := getUserData("Enter birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{ //We instantiate the struct type using the values requested to the user
		firstName: userFirstName, //We assing each value of the custom type with the special operator ":" and a "," at last place
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(), //We can use this function as it is part of the "time" package
	}

	displayUserInput(appUser) //We call the function using only the custom user variable type
}

func displayUserInput(user user) { //We specify on the function that we will receive a "user" type variable
	fmt.Println(user.firstName, user.lastName, user.birthdate, user.createdAt) //We access each variable of the "user"
	//Custom type with the operator "."
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
