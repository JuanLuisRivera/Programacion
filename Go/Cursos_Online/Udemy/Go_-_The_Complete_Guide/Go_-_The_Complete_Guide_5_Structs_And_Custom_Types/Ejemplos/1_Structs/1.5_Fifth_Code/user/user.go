package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Admin struct { //We create a new struct that will be a special case of the "User" struct
	email    string
	password string
	User     //We embed the "User" struct into the "Admin" making the variables of the struct "User" also part of the "Admin" struct
}

func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{ //We hardcode a "User" struct inside the "Admin" struct
			firstName: "Admin",
			lastName:  "Admin",
			birthDate: "-----",
			createdAt: time.Now(),
		},
	}
}

func (user User) DisplayUserInput() {
	fmt.Println(user.firstName, user.lastName, user.birthDate, user.createdAt)
}

func (user *User) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

func New(userFirstName, userLastName, userBirthdate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("first name, last name and birtdate can not be void")
	}

	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}
