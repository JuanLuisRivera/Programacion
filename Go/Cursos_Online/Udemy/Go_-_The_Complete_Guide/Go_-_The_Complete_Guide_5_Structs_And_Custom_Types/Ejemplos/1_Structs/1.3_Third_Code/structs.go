package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func (user user) displayUserInput() { //We specify that this function is a method (It is a function unique to the struct) by using the
	//(user user) notation in front of the "func" keyword (known as receiver argument), that way we specify it will be a "user" method
	fmt.Println(user.firstName, user.lastName, user.birthdate, user.createdAt)
}

func (user *user) clearUserName() { //We create a new function to showcase how methods can manipulate data on the struct
	//It is important to notice that in order to manipulate data we need to use pointers as the system creates copies and if we do not use pointers
	//It will possibly alter only the copy of the struct and not the struct itself, creating unwanted results
	user.firstName = ""
	user.lastName = ""
}

func newUser(userFirstName, userLastName, userBirthdate string) (*user, error) { //We define a constructor for our data type and define an error
	//In case some part of the validation process is incorrect
	if userFirstName == "" || userLastName == "" || userBirthdate == "" { //We validate if the data input is valid
		return nil, errors.New("First name, last name and birtdate can not be void") //We avoid the construction of the variable and display an error
	}

	return &user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}, nil //We create the username and generate an error that shows the process did complete succesfully
}

func main() {
	userFirstName := getUserData("Enter first name: ")
	userLastName := getUserData("Enter last name: ")
	userBirthdate := getUserData("Enter birthdate (MM/DD/YYYY): ")

	appUser, err := newUser(userFirstName, userLastName, userBirthdate) //We create a variable of type "user" called "appUser" using a constructor

	if err != nil { //We validate the data entered by the user and return the error in console
		fmt.Println(err)
		return
	}

	appUser.displayUserInput() //We call the function by specifying it belongs to the "appUser" struct and with the "." operator

	appUser.clearUserName() //By using the function with a pointer we ensure we are modifying the "appUser" struct and not a copy

	appUser.displayUserInput()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value) //Function that specifies the value will be input in only one line
	return value
}
