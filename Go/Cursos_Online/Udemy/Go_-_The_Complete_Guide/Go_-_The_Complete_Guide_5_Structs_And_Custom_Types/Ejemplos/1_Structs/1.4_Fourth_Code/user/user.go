package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct { //We modify the "User" struct to make it possible the use outside the package "user"
	firstName string //We do not modify the variables as we do not want to change the values outside the package but only inside the code
	lastName  string
	birthdate string
	createdAt time.Time
}

func (user User) DisplayUserInput() { //We capitalize the functions to make it possible the use outside the package "user"
	fmt.Println(user.firstName, user.lastName, user.birthdate, user.createdAt)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func New(userFirstName, userLastName, userBirthdate string) (*User, error) { //We modify the function as is no longer necessary to specify it is a
	//"User" package method only, so we can use "New" without confusion on the "main" package
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("first name, last name and birtdate can not be void")
	}

	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}
