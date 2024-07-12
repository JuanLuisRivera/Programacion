package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserData("Enter first name: ")
	userLastName := getUserData("Enter last name: ")
	userBirthdate := getUserData("Enter birthdate (MM/DD/YYYY): ")

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	admin := user.NewAdmin("test@hotmail.com", "test123") //We declare and initialize a new "Admin" variable

	admin.User.DisplayUserInput() //We use the "User" struct methods to display and erase the data in the "Admin" variable

	admin.ClearUserName() //As we did not define a "User" variable in the "admin" struct we can use the methods without having to call ".User"

	admin.User.DisplayUserInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.DisplayUserInput()

	appUser.ClearUserName()

	appUser.DisplayUserInput()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
